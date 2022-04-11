package util

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
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

// ori: Original content. Split with newline
func Input(ori string) string {
	var (
		lines []string
		input []byte
		line  int
		err   error
	)
	sc := bufio.NewScanner(strings.NewReader(ori))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	pri := func() {
		fmt.Println("===========================================")
		for k, v := range lines {
			fmt.Printf("%2d: %s\n", k, v)
		}
		fmt.Println("===========================================")
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		pri()
		fmt.Println("\ne[edit] a[add line] d[delete line] s[swap two line] c[clear] q[quit]")
		input, _, _ = reader.ReadLine()
		switch string(input) {
		case "e", "edit":
			if len(lines) == 0 {
				fmt.Println("No editable rows")
				continue
			}
			fmt.Print("Enter line number: ")
			input, _, err = reader.ReadLine()
			if err != nil {
				fmt.Println("Illegal input")
				continue
			}
			if len(input) == 0 {
				line = len(lines) - 1
			} else {
				line, err = strconv.Atoi(string(input))
				if err != nil {
					fmt.Println("Illegal input")
					continue
				}
				if line < 0 {
					line = 0
				}
				if line >= len(lines) {
					line = len(lines) - 1
				}
			}
			fmt.Printf("Line: [%d]\n", line)
			fmt.Printf("Old: [%s]\n", lines[line])
			fmt.Print("New: ")
			input, _, err = reader.ReadLine()
			if err != nil {
				fmt.Println("Illegal input")
				continue
			}
			fmt.Printf("Input: [%s]\n", string(input))
			lt := string(input)
			fmt.Println("Y/n")
			input, _, err = reader.ReadLine()
			if err != nil {
				fmt.Println("Illegal input")
				continue
			}
			if string(input) == "n" || string(input) == "N" {
				continue
			}
			lines[line] = lt
		case "a", "add":
			fmt.Print("Enter new line number: ")
			input, _, err = reader.ReadLine()
			if err != nil {
				fmt.Println("Illegal input")
				continue
			}
			if len(input) == 0 {
				line = len(lines)
			} else {
				line, err = strconv.Atoi(string(input))
				if err != nil {
					fmt.Println("Illegal input")
					continue
				}
			}
			fmt.Printf("Line number: [%d]\n", line)
			fmt.Println("Y/n")
			input, _, err = reader.ReadLine()
			if err != nil {
				fmt.Println("Illegal input")
				continue
			}
			if string(input) == "n" || string(input) == "N" {
				continue
			}
			if line < 1 {
				lines = append([]string{""}, lines...)
			} else if line >= len(lines) {
				lines = append(lines, "")
			} else {
				lines = append(lines[:line], append([]string{""}, lines[line:]...)...)
			}
		case "d", "delete":
			fmt.Print("Enter line number: ")
			input, _, err = reader.ReadLine()
			if err != nil {
				fmt.Println("Illegal input")
				continue
			}
			line, err = strconv.Atoi(string(input))
			if err != nil {
				fmt.Println("Illegal input")
				continue
			}
			if line < 0 || line >= len(lines) {
				fmt.Println("Illegal input")
				continue
			}
			fmt.Printf("Line: [%d]\n", line)
			fmt.Println("Y/n")
			input, _, err = reader.ReadLine()
			if err != nil {
				fmt.Println("Illegal input")
				continue
			}
			if string(input) == "n" || string(input) == "N" {
				continue
			}
			lines = append(lines[:line], lines[line+1:]...)
		case "s", "swap":
			fmt.Print("Enter line 1 number: ")
			var line1 int
			var line2 int
			// line 1
			input, _, err = reader.ReadLine()
			if err != nil {
				fmt.Println("Illegal input")
				continue
			}
			line1, err = strconv.Atoi(string(input))
			if err != nil {
				fmt.Println("Illegal input")
				continue
			}
			if line1 < 0 || line1 >= len(lines) {
				fmt.Println("Illegal input")
				continue
			}
			// line 2
			fmt.Print("Enter line 2 number: ")
			input, _, err = reader.ReadLine()
			if err != nil {
				fmt.Println("Illegal input")
				continue
			}
			line2, err = strconv.Atoi(string(input))
			if err != nil {
				fmt.Println("Illegal input")
				continue
			}
			if line2 < 0 || line2 >= len(lines) {
				fmt.Println("Illegal input")
				continue
			}
			fmt.Printf("SWAP: [%d] [%d]\n", line1, line2)
			fmt.Println("Y/n")
			input, _, err = reader.ReadLine()
			if err != nil {
				fmt.Println("Illegal input")
				continue
			}
			if string(input) == "n" || string(input) == "N" {
				continue
			}
			lt := lines[line1]
			lines[line1] = lines[line2]
			lines[line2] = lt
		case "c", "clear", "cls":
			fmt.Println("Y/n")
			input, _, err = reader.ReadLine()
			if err != nil {
				fmt.Println("Illegal input")
				continue
			}
			if string(input) == "n" || string(input) == "N" {
				continue
			}
			lines = make([]string, 0)
		case "q", "quit":
			fmt.Println("Your input:")
			pri()
			fmt.Println("Y/n")
			input, _, err = reader.ReadLine()
			if err != nil {
				fmt.Println("Illegal input")
				continue
			}
			if string(input) == "n" || string(input) == "N" {
				continue
			}
			res := strings.TrimSpace(lines[0])
			lines = lines[1:]
			for _, v := range lines {
				res += fmt.Sprintf("\n%s", strings.TrimSpace(v))
			}
			return res
		default:
			fmt.Println("Illegal input")
		}
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

func SplitPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
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
