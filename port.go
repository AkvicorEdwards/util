package util

import (
	"fmt"
	"net"
	"time"
)

func TcpPortIsOpen(ip, port string) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", ip, port), 15*time.Second)
	if err != nil {
		return false
	}
	defer func() { _ = conn.Close() }()
	return true
}

func TcpPortIsOpenByAddr(ipPort string) bool {
	conn, err := net.DialTimeout("tcp", ipPort, 15*time.Second)
	if err != nil {
		return false
	}
	defer func() { _ = conn.Close() }()
	return true
}

func GetAvailablePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}

	defer func(l *net.TCPListener) {
		_ = l.Close()
	}(l)
	return l.Addr().(*net.TCPAddr).Port, nil
}
