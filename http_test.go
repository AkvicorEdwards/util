package util

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRemoteIP(t *testing.T) {
	for _, v := range []struct {
		remoteAddr string
		expected   string
	}{
		{"101.1.0.4:100", "101.1.0.4"},
		{"101.1.0.4:", "101.1.0.4"},
		{"101.1.0.4", ""},
		{":100", ""},
	} {
		if got := RemoteIP(&http.Request{RemoteAddr: v.remoteAddr}); got != v.expected {
			t.Errorf("RemoteAddr:%s expected %s, got %s", v.remoteAddr, v.expected, got)
		}
	}
}

func TestGetClientIP(t *testing.T) {
	r := &http.Request{Header: http.Header{}}
	r.Header.Set("X-Real-IP", " 10.10.10.10  ")
	r.Header.Set("X-Forwarded-For", "  20.20.20.20, 30.30.30.30")
	r.RemoteAddr = "  40.40.40.40:42123 "

	if ip := GetClientIP(r); ip != "20.20.20.20" {
		t.Errorf("expected: 20.20.20.20, got: %s", ip)
	}

	r.Header.Del("X-Forwarded-For")
	if ip := GetClientIP(r); ip != "10.10.10.10" {
		t.Errorf("expected: 10.10.10.10, got: %s", ip)
	}

	r.Header.Set("X-Forwarded-For", "30.30.30.30  ")
	if ip := GetClientIP(r); ip != "30.30.30.30" {
		t.Errorf("expected: 30.30.30.30, got: %s", ip)
	}

	r.Header.Del("X-Forwarded-For")
	r.Header.Del("X-Real-IP")
	if ip := GetClientIP(r); ip != "40.40.40.40" {
		t.Errorf("expected: 40.40.40.40, got: %s", ip)
	}

	r.RemoteAddr = "50.50.50.50"
	if ip := GetClientIP(r); ip != "" {
		t.Errorf("expected: 50.50.50.50, got: %s", ip)
	}
}

func TestGetClientPublicIP(t *testing.T) {
	for _, v := range []struct {
		xForwardedFor string
		remoteAddr    string
		expected      string
	}{
		{"10.3.5.45, 21.45.9.1", "101.1.0.4:100", "21.45.9.1"},
		{"101.3.5.45, 21.45.9.1", "101.1.0.4:100", "101.3.5.45"},
		{"", "101.1.0.4:100", "101.1.0.4"},
		{"21.45.9.1", "101.1.0.4:100", "21.45.9.1"},
		{"21.45.9.1, ", "101.1.0.4:100", "21.45.9.1"},
		{"192.168.5.45, 210.45.9.1, 89.5.6.1", "101.1.0.4:100", "210.45.9.1"},
		{"192.168.5.45, 172.24.9.1, 89.5.6.1", "101.1.0.4:100", "89.5.6.1"},
		{"192.168.5.45, 172.24.9.1", "101.1.0.4:100", "101.1.0.4"},
		{"192.168.5.45, 172.24.9.1", "101.1.0.4:5670", "101.1.0.4"},
	} {
		if got := GetClientPublicIP(&http.Request{
			Header: http.Header{
				"X-Forwarded-For": []string{v.xForwardedFor},
			},
			RemoteAddr: v.remoteAddr,
		}); got != v.expected {
			t.Errorf("IsxForwardedFor:%s, remoteAddr:%s, client ip Should Equal %s", v.xForwardedFor, v.remoteAddr, v.expected)
		}
	}

	r := &http.Request{Header: http.Header{}}
	r.Header.Set("X-Real-IP", " 10.10.10.10  ")
	r.Header.Set("X-Forwarded-For", " 172.17.40.152, 192.168.5.45")
	r.RemoteAddr = "40.40.40.40:42123 "

	if ip := GetClientPublicIP(r); ip != "40.40.40.40" {
		t.Errorf("expected:40.40.40.40, got:%s", ip)
	}

	r.Header.Set("X-Real-IP", " 50.50.50.50  ")
	if ip := GetClientPublicIP(r); ip != "50.50.50.50" {
		t.Errorf("expected:50.50.50.50, got:%s", ip)
	}

	r.Header.Del("X-Real-IP")
	r.Header.Del("X-Forwarded-For")
	r.RemoteAddr = "127.0.0.1:42123 "
	if ip := GetClientPublicIP(r); ip != "" {
		t.Errorf("expected:127.0.0.1, got:%s", ip)
	}
}

func ExampleNewHTTPResp() {
	fmt.Println(NewHTTPResp(HTTPRespCodeOKCode, HTTPRespCodeOKMsg, nil))

	// Output:
	// {"code":0,"msg":"ok"}
}

func ExampleParseHTTPResp() {
	fmt.Println(ParseHTTPResp(`{"code":0,"msg":"ok"}`).String())

	// Output:
	// {"code":0,"msg":"ok"}
}

func ExampleHttpGet() {
	i := NewJSONResult(HttpGet("https://jsonplaceholder.typicode.com/posts/1", nil, HTTPContentTypePlain, nil))
	i1 := NewJSONResult(HttpGet("https://jsonplaceholder.typicode.com/posts", map[string]any{"id": 1}, HTTPContentTypePlain, nil))
	i2 := NewJSONResult(HttpGet("https://jsonplaceholder.typicode.com/posts", map[string]any{"id": 2}, HTTPContentTypePlain, nil))
	if fmt.Sprint(i.Map()["id"]) != "1" {
		fmt.Println(i)
	}
	if fmt.Sprint(i1.Map()["id"]) != "1" {
		fmt.Println(i1)
	}
	if fmt.Sprint(i2.MapArray()[0]["id"]) != "2" {
		fmt.Println(i2)
	}

	// Output:
	//
}

func ExampleHttpPost() {
	i1 := NewJSONResult(HttpPost("https://jsonplaceholder.typicode.com/posts", map[string]any{"title": "t1", "body": "b1", "userId": 1}, HTTPContentTypeUrlencoded, nil))
	if fmt.Sprint(i1.Map()["id"]) != "101" {
		fmt.Println(i1)
	}

	// Output:
	//
}

func ExampleHttpPostGet() {
	i1 := NewJSONResult(HttpPostGet("https://jsonplaceholder.typicode.com/posts", map[string]any{"title": "t2", "body": "b2", "userId": 2}, map[string]any{"title": "t1", "body": "b1", "userId": 1}, HTTPContentTypeUrlencoded, nil))
	if fmt.Sprint(i1.Map()["id"]) != "101" {
		fmt.Println(i1)
	}

	// Output:
	//
}
