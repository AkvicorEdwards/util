package util

import (
  "errors"
  "net"
)

func IsIPAddr(ip string) bool {
  i := net.ParseIP(ip)
  return i != nil
}

func IsLocalIPAddr(ip string) bool {
  i := net.ParseIP(ip)
  if i == nil {
    return false
  }
  return IsLocalIP(i)
}

func IsLocalIP(ip net.IP) bool {
  if ip == nil {
    return false
  }
  if ip.IsLoopback() {
    return true
  }
  
  ip4 := ip.To4()
  if ip4 == nil {
    return false
  }
  
  return ip4[0] == 10 || // 10.0.0.0/8
    (ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) || // 172.16.0.0/12
    (ip4[0] == 169 && ip4[1] == 254) || // 169.254.0.0/16
    (ip4[0] == 192 && ip4[1] == 168) // 192.168.0.0/16
}

var ErrorInvalidIPV4Format = errors.New("invalid ipv4 format")

func IPAddrToUint32(ip string) (uint32, error) {
  i := net.ParseIP(ip)
  if i == nil {
    return 0, ErrorInvalidIPV4Format
  }
  
  b := i.To4()
  if b == nil {
    return 0, ErrorInvalidIPV4Format
  }
  
  return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24, nil
}

func Uint32ToIPAddr(i uint32) string {
  ip := make(net.IP, net.IPv4len)
  ip[0] = byte(i >> 24)
  ip[1] = byte(i >> 16)
  ip[2] = byte(i >> 8)
  ip[3] = byte(i)
  
  return ip.String()
}

func IPToUint32(ip net.IP) (uint32, error) {
  b := ip.To4()
  if b == nil {
    return 0, ErrorInvalidIPV4Format
  }
  
  return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24, nil
}

func Uint32ToIP(i uint32) net.IP {
  ip := make(net.IP, net.IPv4len)
  ip[0] = byte(i >> 24)
  ip[1] = byte(i >> 16)
  ip[2] = byte(i >> 8)
  ip[3] = byte(i)
  
  return ip
}
