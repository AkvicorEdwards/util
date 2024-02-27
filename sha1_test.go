package util

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestSHA1(t *testing.T) {
	// dd if=/dev/zero of=testfile/testFile1 bs=4KB count=4
	// sha1sum testFile1
	const fileSHA1 = "f710c36ffd8c1698f74532968865cf1e8c2d7b3e"
	testString := "Akvicor"
	testStringP1 := "Akv"
	testStringP2 := "icor"
	testSHA1 := "49296886bfedac0dd0399030bc0dbe12e5eed489"
	testSHA1P1 := "0b4e8f1f52b4ae507d2683e760a168bbbd94eef4"
	testBytes := []byte(testString)
	testBytesP1 := []byte(testStringP1)
	testBytesP2 := []byte(testStringP2)
	fileSHA1Bytes, err := hex.DecodeString(fileSHA1)
	if err != nil {
		t.Errorf("fileSHA1Bytes: %v", err)
	}

	// test FromFile & SHA1Result
	s1 := NewSHA1().FromFile("testfile/testFile1")
	if s1.Error() != nil {
		t.Errorf("FromFile failed: %v", s1.Error())
	}
	if s1.Lower() != fileSHA1 {
		t.Errorf("FromFile failed: expected: [%s] got [%s]", fileSHA1, s1.Lower())
	}
	if s1.Upper() != strings.ToUpper(fileSHA1) {
		t.Errorf("FromFile failed: expected: [%s] got [%s]", strings.ToUpper(fileSHA1), s1.Upper())
	}
	if !bytes.Equal(s1.Slice(), fileSHA1Bytes) {
		t.Errorf("FromFile failed: expected: [%v] got [%v]", fileSHA1Bytes, s1.Slice())
	}
	ay := s1.Array()
	if !bytes.Equal(ay[:], fileSHA1Bytes) {
		t.Errorf("FromFile failed: expected: [%v] got [%v]", fileSHA1Bytes, ay)
	}

	// test FromFileChunk
	s1 = NewSHA1().FromFileChunk("testfile/testFile1", 1024)
	if s1.Error() != nil {
		t.Errorf("FromFileChunk failed: %v", s1.Error())
	}
	if s1.Lower() != fileSHA1 {
		t.Errorf("FromFileChunk failed: expected: [%s] got [%s]", fileSHA1, s1.Lower())
	}

	// test FromBytes
	s1 = NewSHA1().FromBytes(testBytes)
	if s1.Error() != nil {
		t.Errorf("FromBytes failed: %v", s1.Error())
	}
	if s1.Lower() != testSHA1 {
		t.Errorf("FromBytes failed: expected: [%s] got [%s]", testSHA1, s1.Lower())
	}

	// test FromString
	s1 = NewSHA1().FromString(testString)
	if s1.Error() != nil {
		t.Errorf("FromString failed: %v", s1.Error())
	}
	if s1.Lower() != testSHA1 {
		t.Errorf("FromString failed: expected: [%s] got [%s]", testSHA1, s1.Lower())
	}

	// test NewSHA1Pip.WriteBytes
	mp := NewSHA1Pip()
	err = mp.WriteBytes(testBytesP1)
	if err != nil {
		t.Errorf("NewSHA1Pip WriteBytes failed: %v", err)
	}
	err = mp.WriteBytes(testBytesP2)
	if err != nil {
		t.Errorf("NewSHA1Pip WriteBytes failed: %v", err)
	}
	s1 = mp.Result()
	if s1.Error() != nil {
		t.Errorf("NewSHA1Pip WriteBytes failed: %v", s1.Error())
	}
	if s1.Lower() != testSHA1 {
		t.Errorf("NewSHA1Pip WriteBytes failed: expected: [%s] got [%s]", testSHA1, s1.Lower())
	}

	// test NewSHA1Pip.WriteString
	mp = NewSHA1Pip()
	err = mp.WriteString(testStringP1)
	if err != nil {
		t.Errorf("NewSHA1Pip WriteString failed: %v", err)
	}
	err = mp.WriteString(testStringP2)
	if err != nil {
		t.Errorf("NewSHA1Pip WriteString failed: %v", err)
	}
	s1 = mp.Result()
	if s1.Error() != nil {
		t.Errorf("NewSHA1Pip WriteString failed: %v", s1.Error())
	}
	if s1.Lower() != testSHA1 {
		t.Errorf("NewSHA1Pip WriteString failed: expected: [%s] got [%s]", testSHA1, s1.Lower())
	}

	// test NewSHA1Pip.Result
	mp = NewSHA1Pip()
	err = mp.WriteString(testStringP1)
	if err != nil {
		t.Errorf("NewSHA1Pip WriteString failed: %v", err)
	}
	s1 = mp.Result()
	if s1.Error() != nil {
		t.Errorf("NewSHA1Pip WriteString failed: %v", s1.Error())
	}
	if s1.Lower() != testSHA1P1 {
		t.Errorf("NewSHA1Pip WriteString failed: expected: [%s] got [%s]", testSHA1P1, s1.Lower())
	}
	err = mp.WriteString(testStringP2)
	if err != nil {
		t.Errorf("NewSHA1Pip WriteString failed: %v", err)
	}
	s1 = mp.Result()
	if s1.Error() != nil {
		t.Errorf("NewSHA1Pip WriteString failed: %v", s1.Error())
	}
	if s1.Lower() != testSHA1 {
		t.Errorf("NewSHA1Pip WriteString failed: expected: [%s] got [%s]", testSHA1, s1.Lower())
	}
}

func ExampleNewSHA1() {
	s := NewSHA1()
	res := s.FromString("Akvicor")
	if res.Error() == nil {
		fmt.Println(res.Lower())
		fmt.Println(res.Upper())
		fmt.Println(res.Slice())
	}

	// Output:
	// 49296886bfedac0dd0399030bc0dbe12e5eed489
	// 49296886BFEDAC0DD0399030BC0DBE12E5EED489
	// [73 41 104 134 191 237 172 13 208 57 144 48 188 13 190 18 229 238 212 137]
}

func ExampleNewSHA1Pip() {
	sp := NewSHA1Pip()
	_, err := io.WriteString(sp, "Akvicor")
	if err != nil {
		fmt.Println(err)
	}
	res := sp.Result()
	if res.Error() == nil {
		fmt.Println(res.Array())
		fmt.Println(res.Slice())
	}

	// Output:
	// [73 41 104 134 191 237 172 13 208 57 144 48 188 13 190 18 229 238 212 137]
	// [73 41 104 134 191 237 172 13 208 57 144 48 188 13 190 18 229 238 212 137]
}
