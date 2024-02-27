package util

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestMD5(t *testing.T) {
	// dd if=/dev/zero of=testfile/testFile1 bs=4KB count=4
	// md5sum testFile1
	const fileMD5 = "1ee0193671609c7d63cfe89b920ad313"
	fileMD5Bytes, err := hex.DecodeString(fileMD5)
	if err != nil {
		t.Errorf("fileMD5Bytes: %v", err)
	}
	testString := "Akvicor"
	testBytes := []byte(testString)
	testStringP1 := "Akv"
	testStringP2 := "icor"
	testBytesP1 := []byte(testStringP1)
	testBytesP2 := []byte(testStringP2)
	testMD5 := "f812705c26adc52561415189e5f78edb"
	testMD5P1 := "d1c13e1e2aa77bed0a694fe1b375a3aa"

	// test FromFile & MD5Result
	m5 := NewMD5().FromFile("testfile/testFile1")
	if m5.Error() != nil {
		t.Errorf("FromFile failed: %v", m5.Error())
	}
	if m5.Lower() != fileMD5 {
		t.Errorf("FromFile failed: expected: [%s] got [%s]", fileMD5, m5.Lower())
	}
	if m5.Upper() != strings.ToUpper(fileMD5) {
		t.Errorf("FromFile failed: expected: [%s] got [%s]", strings.ToUpper(fileMD5), m5.Upper())
	}
	if !bytes.Equal(m5.Slice(), fileMD5Bytes) {
		t.Errorf("FromFile failed: expected: [%v] got [%v]", fileMD5Bytes, m5.Slice())
	}
	ay := m5.Array()
	if !bytes.Equal(ay[:], fileMD5Bytes) {
		t.Errorf("FromFile failed: expected: [%v] got [%v]", fileMD5Bytes, ay)
	}

	// test FromFileChunk
	m5 = NewMD5().FromFileChunk("testfile/testFile1", 1024*64)
	if m5.Error() != nil {
		t.Errorf("FromFileChunk failed: %v", m5.Error())
	}
	if m5.Lower() != fileMD5 {
		t.Errorf("FromFileChunk failed: expected: [%s] got [%s]", fileMD5, m5.Lower())
	}

	// test FromBytes
	m5 = NewMD5().FromBytes(testBytes)
	if m5.Error() != nil {
		t.Errorf("FromBytes failed: %v", m5.Error())
	}
	if m5.Lower() != testMD5 {
		t.Errorf("FromBytes failed: expected: [%s] got [%s]", testMD5, m5.Lower())
	}

	// test FromString
	m5 = NewMD5().FromString(testString)
	if m5.Error() != nil {
		t.Errorf("FromString failed: %v", m5.Error())
	}
	if m5.Lower() != testMD5 {
		t.Errorf("FromString failed: expected: [%s] got [%s]", testMD5, m5.Lower())
	}

	// test NewMD5Pip.WriteBytes
	mp := NewMD5Pip()
	err = mp.WriteBytes(testBytesP1)
	if err != nil {
		t.Errorf("NewMD5Pip WriteBytes failed: %v", err)
	}
	err = mp.WriteBytes(testBytesP2)
	if err != nil {
		t.Errorf("NewMD5Pip WriteBytes failed: %v", err)
	}
	m5 = mp.Result()
	if m5.Error() != nil {
		t.Errorf("NewMD5Pip WriteBytes failed: %v", m5.Error())
	}
	if m5.Lower() != testMD5 {
		t.Errorf("NewMD5Pip WriteBytes failed: expected: [%s] got [%s]", testMD5, m5.Lower())
	}

	// test NewMD5Pip.WriteString
	mp = NewMD5Pip()
	err = mp.WriteString(testStringP1)
	if err != nil {
		t.Errorf("NewMD5Pip WriteString failed: %v", err)
	}
	err = mp.WriteString(testStringP2)
	if err != nil {
		t.Errorf("NewMD5Pip WriteString failed: %v", err)
	}
	m5 = mp.Result()
	if m5.Error() != nil {
		t.Errorf("NewMD5Pip WriteString failed: %v", m5.Error())
	}
	if m5.Lower() != testMD5 {
		t.Errorf("NewMD5Pip WriteString failed: expected: [%s] got [%s]", testMD5, m5.Lower())
	}

	// test NewMD5Pip.Result
	mp = NewMD5Pip()
	err = mp.WriteString(testStringP1)
	if err != nil {
		t.Errorf("NewMD5Pip WriteString failed: %v", err)
	}
	m5 = mp.Result()
	if m5.Error() != nil {
		t.Errorf("NewMD5Pip WriteString failed: %v", m5.Error())
	}
	if m5.Lower() != testMD5P1 {
		t.Errorf("NewMD5Pip WriteString failed: expected: [%s] got [%s]", testMD5P1, m5.Lower())
	}
	err = mp.WriteString(testStringP2)
	if err != nil {
		t.Errorf("NewMD5Pip WriteString failed: %v", err)
	}
	m5 = mp.Result()
	if m5.Error() != nil {
		t.Errorf("NewMD5Pip WriteString failed: %v", m5.Error())
	}
	if m5.Lower() != testMD5 {
		t.Errorf("NewMD5Pip WriteString failed: expected: [%s] got [%s]", testMD5, m5.Lower())
	}
}

func ExampleNewMD5() {
	m5 := NewMD5()
	res := m5.FromString("Akvicor")
	if res.Error() == nil {
		fmt.Println(res.Lower())
		fmt.Println(res.Upper())
		fmt.Println(res.Slice())
	}

	// Output:
	// f812705c26adc52561415189e5f78edb
	// F812705C26ADC52561415189E5F78EDB
	// [248 18 112 92 38 173 197 37 97 65 81 137 229 247 142 219]
}

func ExampleNewMD5Pip() {
	mp := NewMD5Pip()
	_, err := io.WriteString(mp, "Akvicor")
	if err != nil {
		fmt.Println(err)
	}
	res := mp.Result()
	if res.Error() == nil {
		fmt.Println(res.Array())
		fmt.Println(res.Slice())
	}

	// Output:
	// [248 18 112 92 38 173 197 37 97 65 81 137 229 247 142 219]
	// [248 18 112 92 38 173 197 37 97 65 81 137 229 247 142 219]
}
