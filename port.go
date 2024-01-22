package util

import (
  "fmt"
  "net"
  "time"
)

func TcpPortIsOpen(ip, port string) bool {
  conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", ip, port), 3*time.Second)
  if err != nil {
    return false
  }
  defer conn.Close()
  return true
}

func TcpPortIsOpenByAddr(ipPort string) bool {
  conn, err := net.DialTimeout("tcp", ipPort, 3*time.Second)
  if err != nil {
    return false
  }
  defer conn.Close()
  return true
}
