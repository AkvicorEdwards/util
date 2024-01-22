package util

import (
  "bytes"
  "encoding/hex"
  "fmt"
  "io"
  "strings"
  "testing"
)

func TestSHA256(t *testing.T) {
  // dd if=/dev/random of=testfile/testFile1 bs=4KB count=4
  // sha256sum testFile1
  const fileSHA256 = "0aa203abf8927c5a91062311a246ccaaaf038e6bc3a25ea62a978ed5d280698a"
  testString := "Akvicor"
  testStringP1 := "Akv"
  testStringP2 := "icor"
  testSHA256 := "36acc0924a190c77386cd819a3bff60251345e8654399f68e8b740a1afa62ff5"
  testSHA256P1 := "5a01f2f1bd960f10e647fb9c78797836a9774683469ef759c303cd2934548563"
  testBytes := []byte(testString)
  testBytesP1 := []byte(testStringP1)
  testBytesP2 := []byte(testStringP2)
  fileSHA256Bytes, err := hex.DecodeString(fileSHA256)
  if err != nil {
    t.Errorf("fileSHA256Bytes: %v", err)
  }
  
  // test FromFile & SHA256Result
  s1 := NewSHA256().FromFile("testfile/testFile1")
  if s1.Error() != nil {
    t.Errorf("FromFile failed: %v", s1.Error())
  }
  if s1.Lower() != fileSHA256 {
    t.Errorf("FromFile failed: expected: [%s] got [%s]", fileSHA256, s1.Lower())
  }
  if s1.Upper() != strings.ToUpper(fileSHA256) {
    t.Errorf("FromFile failed: expected: [%s] got [%s]", strings.ToUpper(fileSHA256), s1.Upper())
  }
  if !bytes.Equal(s1.Slice(), fileSHA256Bytes) {
    t.Errorf("FromFile failed: expected: [%v] got [%v]", fileSHA256Bytes, s1.Slice())
  }
  ay := s1.Array()
  if !bytes.Equal(ay[:], fileSHA256Bytes) {
    t.Errorf("FromFile failed: expected: [%v] got [%v]", fileSHA256Bytes, ay)
  }
  
  // test FromFileChunk
  s1 = NewSHA256().FromFileChunk("testfile/testFile1", 1024)
  if s1.Error() != nil {
    t.Errorf("FromFileChunk failed: %v", s1.Error())
  }
  if s1.Lower() != fileSHA256 {
    t.Errorf("FromFileChunk failed: expected: [%s] got [%s]", fileSHA256, s1.Lower())
  }
  
  // test FromBytes
  s1 = NewSHA256().FromBytes(testBytes)
  if s1.Error() != nil {
    t.Errorf("FromBytes failed: %v", s1.Error())
  }
  if s1.Lower() != testSHA256 {
    t.Errorf("FromBytes failed: expected: [%s] got [%s]", testSHA256, s1.Lower())
  }
  
  // test FromString
  s1 = NewSHA256().FromString(testString)
  if s1.Error() != nil {
    t.Errorf("FromString failed: %v", s1.Error())
  }
  if s1.Lower() != testSHA256 {
    t.Errorf("FromString failed: expected: [%s] got [%s]", testSHA256, s1.Lower())
  }
  
  // test NewSHA256Pip.WriteBytes
  mp := NewSHA256Pip()
  err = mp.WriteBytes(testBytesP1)
  if err != nil {
    t.Errorf("NewSHA256Pip WriteBytes failed: %v", err)
  }
  err = mp.WriteBytes(testBytesP2)
  if err != nil {
    t.Errorf("NewSHA256Pip WriteBytes failed: %v", err)
  }
  s1 = mp.Result()
  if s1.Error() != nil {
    t.Errorf("NewSHA256Pip WriteBytes failed: %v", s1.Error())
  }
  if s1.Lower() != testSHA256 {
    t.Errorf("NewSHA256Pip WriteBytes failed: expected: [%s] got [%s]", testSHA256, s1.Lower())
  }
  
  // test NewSHA256Pip.WriteString
  mp = NewSHA256Pip()
  err = mp.WriteString(testStringP1)
  if err != nil {
    t.Errorf("NewSHA256Pip WriteString failed: %v", err)
  }
  err = mp.WriteString(testStringP2)
  if err != nil {
    t.Errorf("NewSHA256Pip WriteString failed: %v", err)
  }
  s1 = mp.Result()
  if s1.Error() != nil {
    t.Errorf("NewSHA256Pip WriteString failed: %v", s1.Error())
  }
  if s1.Lower() != testSHA256 {
    t.Errorf("NewSHA256Pip WriteString failed: expected: [%s] got [%s]", testSHA256, s1.Lower())
  }
  
  // test NewSHA256Pip.Result
  mp = NewSHA256Pip()
  err = mp.WriteString(testStringP1)
  if err != nil {
    t.Errorf("NewSHA256Pip WriteString failed: %v", err)
  }
  s1 = mp.Result()
  if s1.Error() != nil {
    t.Errorf("NewSHA256Pip WriteString failed: %v", s1.Error())
  }
  if s1.Lower() != testSHA256P1 {
    t.Errorf("NewSHA256Pip WriteString failed: expected: [%s] got [%s]", testSHA256P1, s1.Lower())
  }
  err = mp.WriteString(testStringP2)
  if err != nil {
    t.Errorf("NewSHA256Pip WriteString failed: %v", err)
  }
  s1 = mp.Result()
  if s1.Error() != nil {
    t.Errorf("NewSHA256Pip WriteString failed: %v", s1.Error())
  }
  if s1.Lower() != testSHA256 {
    t.Errorf("NewSHA256Pip WriteString failed: expected: [%s] got [%s]", testSHA256, s1.Lower())
  }
}

func ExampleNewSHA256() {
  s := NewSHA256()
  res := s.FromString("Akvicor")
  if res.Error() == nil {
    fmt.Println(res.Lower())
    fmt.Println(res.Upper())
    fmt.Println(res.Slice())
  }
  
  // Output:
  // 36acc0924a190c77386cd819a3bff60251345e8654399f68e8b740a1afa62ff5
  // 36ACC0924A190C77386CD819A3BFF60251345E8654399F68E8B740A1AFA62FF5
  // [54 172 192 146 74 25 12 119 56 108 216 25 163 191 246 2 81 52 94 134 84 57 159 104 232 183 64 161 175 166 47 245]
}

func ExampleNewSHA256Pip() {
  sp := NewSHA256Pip()
  io.WriteString(sp, "Akvicor")
  res := sp.Result()
  if res.Error() == nil {
    fmt.Println(res.Array())
    fmt.Println(res.Slice())
  }
  
  // Output:
  // [54 172 192 146 74 25 12 119 56 108 216 25 163 191 246 2 81 52 94 134 84 57 159 104 232 183 64 161 175 166 47 245]
  // [54 172 192 146 74 25 12 119 56 108 216 25 163 191 246 2 81 52 94 134 84 57 159 104 232 183 64 161 175 166 47 245]
}
