package util

import (
	"bytes"
	"encoding/json"
	"errors"
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
		ip2, _, err := net.SplitHostPort(ip)
		if err != nil {
			return ip
		}
		return ip2
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		ip2, _, err := net.SplitHostPort(ip)
		if err != nil {
			return ip
		}
		return ip2
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
			ip2, _, err := net.SplitHostPort(ip)
			if err != nil {
				return ip
			}
			return ip2
		}
	}

	if ip = strings.TrimSpace(r.Header.Get("X-Real-Ip")); ip != "" && !IsLocalIPAddr(ip) {
		ip2, _, err := net.SplitHostPort(ip)
		if err != nil {
			return ip
		}
		return ip2
	}

	if ip = RemoteIP(r); !IsLocalIPAddr(ip) {
		ip2, _, err := net.SplitHostPort(ip)
		if err != nil {
			return ip
		}
		return ip2
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
	Data any          `json:"data,omitempty"`
}

func (r *HTTPRespAPIModel) String() string {
	data, err := json.Marshal(r)
	if err != nil {
		return fmt.Sprintf(`{"code":%d,"msg":"%s"}`, HTTPRespCodeERCode, HTTPRespCodeERMsg)
	}
	return string(data)
}

func (r *HTTPRespAPIModel) Bytes() []byte {
	data, err := json.Marshal(r)
	if err != nil {
		return []byte(fmt.Sprintf(`{"code":%d,"msg":"%s"}`, HTTPRespCodeERCode, HTTPRespCodeERMsg))
	}
	return data
}

// NewHTTPResp Util Reserved code range (-100,100)
func NewHTTPResp(code HTTPRespCode, msg string, data any) *HTTPRespAPIModel {
	return &HTTPRespAPIModel{
		Code: code,
		Msg:  msg,
		Data: data,
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

func WriteHTTPRespAPIOk(w http.ResponseWriter, data any, msg ...any) {
	m := HTTPRespCodeOKMsg
	if len(msg) > 0 {
		m = fmt.Sprint(msg...)
	}
	w.Header().Set("Content-Type", HTTPContentTypeJson)
	_, _ = w.Write(NewHTTPResp(HTTPRespCodeOKCode, m, data).Bytes())
}

func WriteHTTPRespAPIFailed(w http.ResponseWriter, data any, msg ...any) {
	m := HTTPRespCodeERMsg
	if len(msg) > 0 {
		m = fmt.Sprint(msg...)
	}
	w.Header().Set("Content-Type", HTTPContentTypeJson)
	_, _ = w.Write(NewHTTPResp(HTTPRespCodeERCode, m, data).Bytes())
}

func WriteHTTPRespAPIInvalidKey(w http.ResponseWriter, data any, msg ...any) {
	m := HTTPRespCodeInvalidKeyMsg
	if len(msg) > 0 {
		m = fmt.Sprint(msg...)
	}
	w.Header().Set("Content-Type", HTTPContentTypeJson)
	_, _ = w.Write(NewHTTPResp(HTTPRespCodeInvalidKeyCode, m, data).Bytes())
}

func WriteHTTPRespAPIInvalidInput(w http.ResponseWriter, data any, msg ...any) {
	m := HTTPRespCodeInvalidInputMsg
	if len(msg) > 0 {
		m = fmt.Sprint(msg...)
	}
	w.Header().Set("Content-Type", HTTPContentTypeJson)
	_, _ = w.Write(NewHTTPResp(HTTPRespCodeInvalidInputCode, m, data).Bytes())
}

func WriteHTTPRespAPIProcessingFailed(w http.ResponseWriter, data any, msg ...any) {
	m := HTTPRespCodeProcessingFailedMsg
	if len(msg) > 0 {
		m = fmt.Sprint(msg...)
	}
	w.Header().Set("Content-Type", HTTPContentTypeJson)
	_, _ = w.Write(NewHTTPResp(HTTPRespCodeProcessingFailedCode, m, data).Bytes())
}

const (
	HTTPContentTypeUrlencoded     = "application/x-www-form-urlencoded"
	HTTPContentTypeUrlencodedUTF8 = "application/x-www-form-urlencoded; charset=utf-8"
	HTTPContentTypeJson           = "application/json"
	HTTPContentTypeJsonUTF8       = "application/json; charset=utf-8"
	HTTPContentTypeXml            = "application/xml"
	HTTPContentTypeXmlUTF8        = "application/xml; charset=utf-8"
	HTTPContentTypePlain          = "text/plain"
	HTTPContentTypePlainUTF8      = "text/plain; charset=utf-8"
	HTTPContentTypeHtml           = "text/html"
	HTTPContentTypeHtmlUTF8       = "text/html; charset=utf-8"
	HTTPContentTypeFormData       = "multipart/form-data"
	HTTPContentTypeFormDataUTF8   = "multipart/form-data; charset=utf-8"
)

func HttpGet(u string, args any, contentType string, header map[string]string) ([]byte, error) {
	arg := NewJSON(args, false).Map()
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	for k, v := range arg {
		q.Add(k, fmt.Sprint(v))
	}
	req.URL.RawQuery = q.Encode()
	if header == nil {
		header = make(map[string]string)
	}
	if len(contentType) > 0 {
		header["Content-Type"] = contentType
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func HttpPost(u string, args any, contentType string, header map[string]string) ([]byte, error) {
	var req *http.Request
	var err error

	if contentType == HTTPContentTypeUrlencoded || contentType == HTTPContentTypeUrlencodedUTF8 {
		arg := NewJSON(args, false).Map()
		payload := url.Values{}
		for k, v := range arg {
			payload.Set(k, fmt.Sprint(v))
		}
		req, err = http.NewRequest("POST", u, strings.NewReader(payload.Encode()))
		if err != nil {
			return nil, err
		}
	} else if contentType == HTTPContentTypeJson || contentType == HTTPContentTypeJsonUTF8 {
		req, err = http.NewRequest("POST", u, bytes.NewBuffer(NewJSON(args, false).Bytes()))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("unknown content type")
	}
	if header == nil {
		header = make(map[string]string)
	}
	if len(contentType) > 0 {
		header["Content-Type"] = contentType
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func HttpPostGet(u string, argsGET, argsPOST any, contentType string, header map[string]string) ([]byte, error) {
	var req *http.Request
	var err error

	if contentType == HTTPContentTypeUrlencoded || contentType == HTTPContentTypeUrlencodedUTF8 {
		arg := NewJSON(argsPOST, false).Map()
		payload := url.Values{}
		for k, v := range arg {
			payload.Set(k, fmt.Sprint(v))
		}
		req, err = http.NewRequest("POST", u, strings.NewReader(payload.Encode()))
		if err != nil {
			return nil, err
		}
	} else if contentType == HTTPContentTypeJson || contentType == HTTPContentTypeJsonUTF8 {
		req, err = http.NewRequest("POST", u, bytes.NewBuffer(NewJSON(argsPOST, false).Bytes()))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("unknown content type")
	}
	if header == nil {
		header = make(map[string]string)
	}
	if len(contentType) > 0 {
		header["Content-Type"] = contentType
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}

	argGet := NewJSON(argsGET, false).Map()
	q := req.URL.Query()
	for k, v := range argGet {
		q.Add(k, fmt.Sprint(v))
	}
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
