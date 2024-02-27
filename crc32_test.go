package util

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestCRC32(t *testing.T) {
	// dd if=/dev/zero of=testfile/testFile1 bs=4KB count=4
	// crc32 testFile1
	const fileCRC32 = "577ce78d"
	fileCRC32Bytes, err := hex.DecodeString(fileCRC32)
	if err != nil {
		t.Errorf("fileCRC32Bytes: %v", err)
	}
	testString := "Akvicor"
	testBytes := []byte(testString)
	testStringP1 := "Akv"
	testStringP2 := "icor"
	testBytesP1 := []byte(testStringP1)
	testBytesP2 := []byte(testStringP2)
	testCRC32 := "3b60e1cd"
	testCRC32P1 := "b1765880"

	// test FromFile & CRC32Result
	crc := NewCRC32().FromFile("testfile/testFile1")
	if crc.Error() != nil {
		t.Errorf("FromFile failed: %v", crc.Error())
	}
	if crc.Value() != BytesSliceToUInt32(fileCRC32Bytes) {
		t.Errorf("FromFile failed: expected: [%d] got [%d]", BytesSliceToUInt32(fileCRC32Bytes), crc.Value())
	}
	if crc.Lower() != fileCRC32 {
		t.Errorf("FromFile failed: expected: [%s] got [%s]", fileCRC32, crc.Lower())
	}
	if crc.Upper() != strings.ToUpper(fileCRC32) {
		t.Errorf("FromFile failed: expected: [%s] got [%s]", strings.ToUpper(fileCRC32), crc.Upper())
	}
	if !bytes.Equal(crc.Slice(), fileCRC32Bytes) {
		t.Errorf("FromFile failed: expected: [%v] got [%v]", fileCRC32Bytes, crc.Slice())
	}
	ay := crc.Array()
	if !bytes.Equal(ay[:], fileCRC32Bytes) {
		t.Errorf("FromFile failed: expected: [%v] got [%v]", fileCRC32Bytes, ay)
	}

	// test FromFileChunk
	crc = NewCRC32().FromFileChunk("testfile/testFile1", 1024*64)
	if crc.Error() != nil {
		t.Errorf("FromFileChunk failed: %v", crc.Error())
	}
	if crc.Lower() != fileCRC32 {
		t.Errorf("FromFileChunk failed: expected: [%s] got [%s]", fileCRC32, crc.Lower())
	}

	// test FromBytes
	crc = NewCRC32().FromBytes(testBytes)
	if crc.Error() != nil {
		t.Errorf("FromBytes failed: %v", crc.Error())
	}
	if crc.Lower() != testCRC32 {
		t.Errorf("FromBytes failed: expected: [%s] got [%s]", testCRC32, crc.Lower())
	}

	// test FromString
	crc = NewCRC32().FromString(testString)
	if crc.Error() != nil {
		t.Errorf("FromString failed: %v", crc.Error())
	}
	if crc.Lower() != testCRC32 {
		t.Errorf("FromString failed: expected: [%s] got [%s]", testCRC32, crc.Lower())
	}

	// test NewCRC32Pip.Write
	cp := NewCRC32Pip()
	n, err := cp.Write(testBytesP1)
	if n != len(testBytesP1) || err != nil {
		t.Errorf("NewCRC32Pip Write failed: [%d][%d] %v", n, len(testBytesP1), err)
	}
	n, err = cp.Write(testBytesP2)
	if n != len(testBytesP2) || err != nil {
		t.Errorf("NewCRC32Pip Write failed: [%d][%d] %v", n, len(testBytesP2), err)
	}
	crc = cp.Result()
	if crc.Error() != nil {
		t.Errorf("NewCRC32Pip Write failed: %v", crc.Error())
	}
	if crc.Lower() != testCRC32 {
		t.Errorf("NewCRC32Pip Write failed: expected: [%s] got [%s]", testCRC32, crc.Lower())
	}

	// test NewCRC32Pip.WriteBytes
	cp = NewCRC32Pip()
	cp.WriteBytes(testBytesP1)
	cp.WriteBytes(testBytesP2)
	crc = cp.Result()
	if crc.Error() != nil {
		t.Errorf("NewCRC32Pip WriteBytes failed: %v", crc.Error())
	}
	if crc.Lower() != testCRC32 {
		t.Errorf("NewCRC32Pip WriteBytes failed: expected: [%s] got [%s]", testCRC32, crc.Lower())
	}

	// test NewCRC32Pip.WriteString
	cp = NewCRC32Pip()
	cp.WriteString(testStringP1)
	cp.WriteString(testStringP2)
	crc = cp.Result()
	if crc.Error() != nil {
		t.Errorf("NewCRC32Pip WriteString failed: %v", crc.Error())
	}
	if crc.Lower() != testCRC32 {
		t.Errorf("NewCRC32Pip WriteString failed: expected: [%s] got [%s]", testCRC32, crc.Lower())
	}

	// test NewCRC32Pip.Result
	cp = NewCRC32Pip()
	cp.WriteString(testStringP1)
	crc = cp.Result()
	if crc.Error() != nil {
		t.Errorf("NewCRC32Pip WriteString failed: %v", crc.Error())
	}
	if crc.Lower() != testCRC32P1 {
		t.Errorf("NewCRC32Pip WriteString failed: expected: [%s] got [%s]", testCRC32P1, crc.Lower())
	}
	cp.WriteString(testStringP2)
	crc = cp.Result()
	if crc.Error() != nil {
		t.Errorf("NewCRC32Pip WriteString failed: %v", crc.Error())
	}
	if crc.Lower() != testCRC32 {
		t.Errorf("NewCRC32Pip WriteString failed: expected: [%s] got [%s]", testCRC32, crc.Lower())
	}
}

func ExampleNewCRC32() {
	crc := NewCRC32()
	res := crc.FromString("Akvicor")
	if res.Error() == nil {
		fmt.Println(res.Value())
		fmt.Println(res.Lower())
		fmt.Println(res.Upper())
	}

	// Output:
	// 996205005
	// 3b60e1cd
	// 3B60E1CD
}

func ExampleNewCRC32Pip() {
	cp := NewCRC32Pip()
	_, err := io.WriteString(cp, "Akvicor")
	if err != nil {
		fmt.Println(err)
	}
	res := cp.Result()
	if res.Error() == nil {
		fmt.Println(res.Array())
		fmt.Println(res.Slice())
	}
	// Output:
	// [59 96 225 205]
	// [59 96 225 205]
}
