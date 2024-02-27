package util

import (
	"fmt"
	"net"
	"testing"
)

func TestIsIPAddr(t *testing.T) {
	for ip, expected := range map[string]bool{
		"":                   false,
		"invalid ip address": false,
		"127.0.0.1":          true,
		"::1":                true,
		"1.1.1.1":            true,
	} {
		if IsIPAddr(ip) != expected {
			t.Errorf("ip %s expected %v, got %v", ip, expected, IsIPAddr(ip))
		}
	}
}

func TestIsLocalIPAddr(t *testing.T) {
	for ip, expected := range map[string]bool{
		"":                   false,
		"invalid ip address": false,
		"127.0.0.1":          true,
		"::1":                true,
		"182.56.9.18":        false,
		"192.168.9.18":       true,
		"10.168.9.18":        true,
		"11.168.9.18":        false,
		"172.16.9.18":        true,
		"172.17.9.18":        true,
		"172.18.9.18":        true,
		"172.19.9.18":        true,
		"172.20.9.18":        true,
		"172.21.9.18":        true,
		"172.22.9.18":        true,
		"172.23.9.18":        true,
		"172.24.9.18":        true,
		"172.25.9.18":        true,
		"172.26.9.18":        true,
		"172.27.9.18":        true,
		"172.28.9.18":        true,
		"172.29.9.18":        true,
		"172.30.9.18":        true,
		"172.31.9.18":        true,
		"172.32.9.18":        false,
	} {
		if IsLocalIPAddr(ip) != expected {
			t.Errorf("ip %s expected %v, got %v", ip, expected, IsLocalIPAddr(ip))
		}
	}
}

func TestIsLocalIP(t *testing.T) {
	for ip, expected := range map[string]bool{
		"":                   false,
		"invalid ip address": false,
		"127.0.0.1":          true,
		"::1":                true,
		"182.56.9.18":        false,
		"192.168.9.18":       true,
		"10.168.9.18":        true,
		"11.168.9.18":        false,
		"172.16.9.18":        true,
		"172.17.9.18":        true,
		"172.18.9.18":        true,
		"172.19.9.18":        true,
		"172.20.9.18":        true,
		"172.21.9.18":        true,
		"172.22.9.18":        true,
		"172.23.9.18":        true,
		"172.24.9.18":        true,
		"172.25.9.18":        true,
		"172.26.9.18":        true,
		"172.27.9.18":        true,
		"172.28.9.18":        true,
		"172.29.9.18":        true,
		"172.30.9.18":        true,
		"172.31.9.18":        true,
		"172.32.9.18":        false,
	} {
		if IsLocalIP(net.ParseIP(ip)) != expected {
			t.Errorf("ip %s expected %v, got %v", ip, expected, IsLocalIP(net.ParseIP(ip)))
		}
	}
}

func TestIPAddrToUint32(t *testing.T) {
	for _, v := range []struct {
		ip  string
		val uint32
	}{
		{"127.0.0.1", 2130706433},
		{"0.0.0.0", 0},
		{"255.255.255.255", 4294967295},
		{"192.168.1.1", 3232235777},
	} {
		got, err := IPAddrToUint32(v.ip)
		if err != nil {
			t.Errorf("ip:%s expected %d, got:%d err:%v", v.ip, v.val, got, err)
		}

		if got != v.val {
			t.Errorf("ip:%s expected %d, got:%d", v.ip, v.val, got)
		}
	}

	for _, ip := range []string{
		"",
		"invalid ip address",
		"::1",
	} {
		_, err := IPAddrToUint32(ip)
		if err == nil {
			t.Errorf("ip:%s invalid IP passes", ip)
		}
	}
}

func TestUint32ToIPAddr(t *testing.T) {
	for _, v := range []struct {
		ip  string
		val uint32
	}{
		{"127.0.0.1", 2130706433},
		{"0.0.0.0", 0},
		{"255.255.255.255", 4294967295},
		{"192.168.1.1", 3232235777},
	} {
		got := Uint32ToIPAddr(v.val)

		if got != v.ip {
			t.Errorf("val: %d, expected:%s, got:%s", v.val, v.ip, got)
		}
	}
}

func TestIPToUint32(t *testing.T) {
	for _, v := range []struct {
		ip  string
		val uint32
	}{
		{"127.0.0.1", 2130706433},
		{"0.0.0.0", 0},
		{"255.255.255.255", 4294967295},
		{"192.168.1.1", 3232235777},
	} {
		got, err := IPToUint32(net.ParseIP(v.ip))
		if err != nil {
			t.Errorf("ip:%s expected %d, got:%d err:%v", v.ip, v.val, got, err)
		}

		if got != v.val {
			t.Errorf("ip:%s expected %d, got:%d", v.ip, v.val, got)
		}
	}

	for _, ip := range []string{
		"",
		"invalid ip address",
		"::1",
	} {
		_, err := IPToUint32(net.ParseIP(ip))
		if err == nil {
			t.Errorf("ip:%s invalid IP passes", ip)
		}
	}
}

func TestUint32ToIP(t *testing.T) {
	for _, v := range []struct {
		ip  string
		val uint32
	}{
		{"127.0.0.1", 2130706433},
		{"0.0.0.0", 0},
		{"255.255.255.255", 4294967295},
		{"192.168.1.1", 3232235777},
	} {
		got := Uint32ToIP(v.val)

		if got.String() != v.ip {
			t.Errorf("val: %d, expected:%s, got:%s", v.val, v.ip, got)
		}
	}
}

func ExampleGetLocalIp() {
	p, err := GetLocalIp()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
}
