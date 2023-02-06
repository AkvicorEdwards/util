package util

import (
	"bytes"
	"math/rand"
	"net"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}

func BitSet(b interface{}, bit byte, set bool) {
	switch b.(type) {
	case *byte:
		if set {
			*b.(*byte) = *b.(*byte) | (byte(1) << bit)
		} else {
			*b.(*byte) = *b.(*byte) & (^(byte(1) << bit))
		}
	case *uint16:
		if set {
			*b.(*uint16) = *b.(*uint16) | (uint16(1) << bit)
		} else {
			*b.(*uint16) = *b.(*uint16) & (^(uint16(1) << bit))
		}
	case *uint32:
		if set {
			*b.(*uint32) = *b.(*uint32) | (uint32(1) << bit)
		} else {
			*b.(*uint32) = *b.(*uint32) & (^(uint32(1) << bit))
		}
	case *uint64:
		if set {
			*b.(*uint64) = *b.(*uint64) | (uint64(1) << bit)
		} else {
			*b.(*uint64) = *b.(*uint64) & (^(uint64(1) << bit))
		}
	case *uint:
		if set {
			*b.(*uint) = *b.(*uint) | (uint(1) << bit)
		} else {
			*b.(*uint) = *b.(*uint) & (^(uint(1) << bit))
		}
	default:
		b = nil
	}
}

//	0: File not exist
//	1: File is directory
//	2: File is file
func FileStat(filename string) int {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return 0
	}
	if info.IsDir() {
		return 1
	} else {
		return 2
	}
}

func SplitPathSkip(p string, skip int) (head, tail string) {
	p = path.Clean("/" + p + "/")
	fin := 0
	cur := 0
	for skip >= 0 {
		skip--
		if fin+1 >= len(p) {
			return p, ""
		}
		cur = strings.IndexByte(p[fin+1:], '/')
		if cur < 0 {
			return p, ""
		}
		fin += cur + 1
	}
	return p[:fin], p[fin:]
}

func SplitPath(p string) (head, tail string) {
	p = path.Clean("/" + p + "/")
	return SplitPathSkip(p, 0)
}

func SplitPathRepeat(p string, repeat int) (head, tail string) {
	p = path.Clean("/" + p + "/")
	head, tail = SplitPathSkip(p, repeat)
	c := strings.Count(head, "/") - 1
	if c == repeat {
		i := strings.LastIndexByte(head, '/')
		return head[i:], tail
	} else if c < repeat {
		return "/", ""
	}
	return head, tail
}

func RandomString(length int, str ...string) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	if len(str) != 0 {
		chars = ""
		for _, v := range str {
			chars += v
		}
	}
	charsLen := len(chars)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	res := make([]byte, length)
	for i := range res {
		res[i] = chars[r.Intn(charsLen)]
	}
	return string(res)
}

func RandomStringWithTimestamp(length int) string {
	if length < 7 {
		length = 7
	}
	nowStampStr := strconv.FormatInt(time.Now().Unix(), 36)
	fill := make([]byte, 7-len(nowStampStr))
	for i := range fill {
		fill[i] = '0'
	}
	return string(fill) + nowStampStr + RandomString(length-7)
}

func ParseRandomStringWithTimestamp(str string) (int64, string) {
	if len(str) < 7 {
		return 0, ""
	}
	date, err := strconv.ParseInt(string([]byte(str)[:7]), 36, 64)
	if err != nil {
		return 0, ""
	}
	return date, string([]byte(str)[7:])
}

func GetIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if net.ParseIP(ip) != nil {
		return ip
	}

	ip = r.Header.Get("X-Forward-For")
	for _, i := range strings.Split(ip, ",") {
		if net.ParseIP(i) != nil {
			return i
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "unknown"
	}

	if net.ParseIP(ip) != nil {
		return ip
	}

	return "unknown"
}
