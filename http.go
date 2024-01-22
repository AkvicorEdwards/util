package util

import (
  "bytes"
  "encoding/json"
  "fmt"
  "io"
  "net"
  "net/http"
  "net/url"
  "strings"
)

func RemoteIP(r *http.Request) string {
  ip, _, _ := net.SplitHostPort(r.RemoteAddr)
  return ip
}

func GetClientIP(r *http.Request) string {
  ip := strings.TrimSpace(strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0])
  if ip != "" {
    return ip
  }
  
  ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
  if ip != "" {
    return ip
  }
  
  if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
    return ip
  }
  
  return ""
}

func GetClientPublicIP(r *http.Request) string {
  var ip string
  for _, ip = range strings.Split(r.Header.Get("X-Forwarded-For"), ",") {
    if ip = strings.TrimSpace(ip); ip != "" && !IsLocalIPAddr(ip) {
      return ip
    }
  }
  
  if ip = strings.TrimSpace(r.Header.Get("X-Real-Ip")); ip != "" && !IsLocalIPAddr(ip) {
    return ip
  }
  
  if ip = RemoteIP(r); !IsLocalIPAddr(ip) {
    return ip
  }
  
  return ""
}

func RespRedirect(w http.ResponseWriter, r *http.Request, url string) {
  http.Redirect(w, r, url, http.StatusMovedPermanently)
}

func LastPage(w http.ResponseWriter, r *http.Request) {
  RespRedirect(w, r, r.URL.String()[:strings.LastIndexByte(r.URL.String(), '/')])
}

func Reload(w http.ResponseWriter, r *http.Request) {
  RespRedirect(w, r, r.URL.String())
}

// HTTPRespCode Util Reserved code range (-100,100)
type HTTPRespCode int

const (
  HTTPRespCodeOKCode               HTTPRespCode = 0
  HTTPRespCodeOKMsg                string       = "ok"
  HTTPRespCodeERCode               HTTPRespCode = 1
  HTTPRespCodeERMsg                string       = "failed to respond"
  HTTPRespCodeInvalidKeyCode       HTTPRespCode = 2
  HTTPRespCodeInvalidKeyMsg        string       = "invalid key"
  HTTPRespCodeProcessingFailedCode HTTPRespCode = 3
  HTTPRespCodeProcessingFailedMsg  string       = "processing failed"
  HTTPRespCodeInvalidInputCode     HTTPRespCode = 4
  HTTPRespCodeInvalidInputMsg      string       = "invalid input"
)

type HTTPRespAPIModel struct {
  Code HTTPRespCode `json:"code"`
  Msg  string       `json:"msg"`
}

func (r *HTTPRespAPIModel) String() string {
  data, err := json.Marshal(r)
  if err != nil {
    return fmt.Sprintf(`{"code":%d,"msg":"%s"}`, HTTPRespCodeERCode, HTTPRespCodeERMsg)
  }
  return string(data)
}

// NewHTTPResp Util Reserved code range (-100,100)
func NewHTTPResp(code HTTPRespCode, msg string) *HTTPRespAPIModel {
  return &HTTPRespAPIModel{
    Code: code,
    Msg:  msg,
  }
}

func ParseHTTPResp(respstr string) *HTTPRespAPIModel {
  resp := &HTTPRespAPIModel{}
  err := json.Unmarshal([]byte(respstr), resp)
  if err != nil {
    return nil
  }
  return resp
}

func WriteHTTPRespAPIOk(w http.ResponseWriter, msg ...any) {
  m := HTTPRespCodeOKMsg
  if len(msg) > 0 {
    m = fmt.Sprint(msg...)
  }
  _, _ = w.Write([]byte(NewHTTPResp(HTTPRespCodeOKCode, m).String()))
}

func WriteHTTPRespAPIFailed(w http.ResponseWriter, msg ...any) {
  m := HTTPRespCodeERMsg
  if len(msg) > 0 {
    m = fmt.Sprint(msg...)
  }
  _, _ = w.Write([]byte(NewHTTPResp(HTTPRespCodeERCode, m).String()))
}

func WriteHTTPRespAPIInvalidKey(w http.ResponseWriter, msg ...any) {
  m := HTTPRespCodeInvalidKeyMsg
  if len(msg) > 0 {
    m = fmt.Sprint(msg...)
  }
  _, _ = w.Write([]byte(NewHTTPResp(HTTPRespCodeInvalidKeyCode, m).String()))
}

func WriteHTTPRespAPIInvalidInput(w http.ResponseWriter, msg ...any) {
  m := HTTPRespCodeInvalidInputMsg
  if len(msg) > 0 {
    m = fmt.Sprint(msg...)
  }
  _, _ = w.Write([]byte(NewHTTPResp(HTTPRespCodeInvalidInputCode, m).String()))
}

func WriteHTTPRespAPIProcessingFailed(w http.ResponseWriter, msg ...any) {
  m := HTTPRespCodeProcessingFailedMsg
  if len(msg) > 0 {
    m = fmt.Sprint(msg...)
  }
  _, _ = w.Write([]byte(NewHTTPResp(HTTPRespCodeProcessingFailedCode, m).String()))
}

const (
  HTTPContentTypeUrlencoded = "application/x-www-form-urlencoded; charset=utf-8"
  HTTPContentTypeJson       = "application/json; charset=UTF-8"
  // HTTPContentTypePlain      = "text/plain; charset=utf-8"
  // HTTPContentTypeFormData   = "multipart/form-data; charset=utf-8"
)

func HttpGet(u string, args any) []byte {
  arg := NewJSON(args, false).Map()
  req, err := http.NewRequest("GET", u, nil)
  if err != nil {
    return nil
  }
  q := req.URL.Query()
  for k, v := range arg {
    q.Add(k, fmt.Sprint(v))
  }
  req.URL.RawQuery = q.Encode()
  req.Header.Add("Content-Type", HTTPContentTypeJson)
  
  client := &http.Client{}
  rsp, err := client.Do(req)
  if err != nil {
    return nil
  }
  defer rsp.Body.Close()
  
  body, err := io.ReadAll(rsp.Body)
  if err != nil {
    return nil
  }
  return body
}

func HttpPost(contentType string, u string, args any) []byte {
  var req *http.Request
  var err error
  
  if contentType == HTTPContentTypeUrlencoded {
    arg := NewJSON(args, false).Map()
    payload := url.Values{}
    for k, v := range arg {
      payload.Set(k, fmt.Sprint(v))
    }
    req, err = http.NewRequest("POST", u, strings.NewReader(payload.Encode()))
    if err != nil {
      return nil
    }
  } else if contentType == HTTPContentTypeJson {
    req, err = http.NewRequest("POST", u, bytes.NewBuffer(NewJSON(args, false).Bytes()))
    if err != nil {
      return nil
    }
  } else {
    return nil
  }
  req.Header.Add("Content-Type", contentType)
  
  client := &http.Client{}
  rsp, err := client.Do(req)
  if err != nil {
    return nil
  }
  defer rsp.Body.Close()
  
  body, err := io.ReadAll(rsp.Body)
  if err != nil {
    return nil
  }
  return body
}

func HttpPostGet(contentType string, u string, argsGET, argsPOST any) []byte {
  var req *http.Request
  var err error
  
  if contentType == HTTPContentTypeUrlencoded {
    arg := NewJSON(argsPOST, false).Map()
    payload := url.Values{}
    for k, v := range arg {
      payload.Set(k, fmt.Sprint(v))
    }
    req, err = http.NewRequest("POST", u, strings.NewReader(payload.Encode()))
    if err != nil {
      return nil
    }
  } else if contentType == HTTPContentTypeJson {
    req, err = http.NewRequest("POST", u, bytes.NewBuffer(NewJSON(argsPOST, false).Bytes()))
    if err != nil {
      return nil
    }
  } else {
    return nil
  }
  req.Header.Add("Content-Type", contentType)
  
  argGet := NewJSON(argsGET, false).Map()
  q := req.URL.Query()
  for k, v := range argGet {
    q.Add(k, fmt.Sprint(v))
  }
  req.URL.RawQuery = q.Encode()
  
  client := &http.Client{}
  rsp, err := client.Do(req)
  if err != nil {
    return nil
  }
  defer rsp.Body.Close()
  
  body, err := io.ReadAll(rsp.Body)
  if err != nil {
    return nil
  }
  return body
}
