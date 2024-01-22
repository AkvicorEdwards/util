package util

import (
  "bytes"
  "encoding/hex"
  "fmt"
  "io"
  "strings"
  "testing"
)

func TestSHA512(t *testing.T) {
  // dd if=/dev/random of=testfile/testFile1 bs=4KB count=4
  // sha512sum testFile1
  const fileSHA512 = "9b5ca4ca0bd1c835bbf2cb9d9fa0de95107aea76735315fdf7628f30428fbbe7f4254029bb2d6229b7b035e732d4f6582a1a39eb5ae6820389e6de5f6e0f14e5"
  testString := "Akvicor"
  testStringP1 := "Akv"
  testStringP2 := "icor"
  testSHA512 := "6fdfb481af6573ac49d4031d22f0d73f0f24ea8cb113982593ec65ab085b0635b46d7683aa3f4ef36c2cb25a7bb8ba8cbae2fa3e9810d35b1c70a93f37586361"
  testSHA512P1 := "3c7ec9519b880b2a6438dea41c4d49cb8761f167dadb237690799aae202852c688d7234f34fa8ff468b794cbfff542e5bd9c77a20f74db7aff7f0b716b49af9c"
  testBytes := []byte(testString)
  testBytesP1 := []byte(testStringP1)
  testBytesP2 := []byte(testStringP2)
  fileSHA512Bytes, err := hex.DecodeString(fileSHA512)
  if err != nil {
    t.Errorf("fileSHA512Bytes: %v", err)
  }
  
  // test FromFile & SHA512Result
  s1 := NewSHA512().FromFile("testfile/testFile1")
  if s1.Error() != nil {
    t.Errorf("FromFile failed: %v", s1.Error())
  }
  if s1.Lower() != fileSHA512 {
    t.Errorf("FromFile failed: expected: [%s] got [%s]", fileSHA512, s1.Lower())
  }
  if s1.Upper() != strings.ToUpper(fileSHA512) {
    t.Errorf("FromFile failed: expected: [%s] got [%s]", strings.ToUpper(fileSHA512), s1.Upper())
  }
  if !bytes.Equal(s1.Slice(), fileSHA512Bytes) {
    t.Errorf("FromFile failed: expected: [%v] got [%v]", fileSHA512Bytes, s1.Slice())
  }
  ay := s1.Array()
  if !bytes.Equal(ay[:], fileSHA512Bytes) {
    t.Errorf("FromFile failed: expected: [%v] got [%v]", fileSHA512Bytes, ay)
  }
  
  // test FromFileChunk
  s1 = NewSHA512().FromFileChunk("testfile/testFile1", 1024)
  if s1.Error() != nil {
    t.Errorf("FromFileChunk failed: %v", s1.Error())
  }
  if s1.Lower() != fileSHA512 {
    t.Errorf("FromFileChunk failed: expected: [%s] got [%s]", fileSHA512, s1.Lower())
  }
  
  // test FromBytes
  s1 = NewSHA512().FromBytes(testBytes)
  if s1.Error() != nil {
    t.Errorf("FromBytes failed: %v", s1.Error())
  }
  if s1.Lower() != testSHA512 {
    t.Errorf("FromBytes failed: expected: [%s] got [%s]", testSHA512, s1.Lower())
  }
  
  // test FromString
  s1 = NewSHA512().FromString(testString)
  if s1.Error() != nil {
    t.Errorf("FromString failed: %v", s1.Error())
  }
  if s1.Lower() != testSHA512 {
    t.Errorf("FromString failed: expected: [%s] got [%s]", testSHA512, s1.Lower())
  }
  
  // test NewSHA512Pip.WriteBytes
  mp := NewSHA512Pip()
  err = mp.WriteBytes(testBytesP1)
  if err != nil {
    t.Errorf("NewSHA512Pip WriteBytes failed: %v", err)
  }
  err = mp.WriteBytes(testBytesP2)
  if err != nil {
    t.Errorf("NewSHA512Pip WriteBytes failed: %v", err)
  }
  s1 = mp.Result()
  if s1.Error() != nil {
    t.Errorf("NewSHA512Pip WriteBytes failed: %v", s1.Error())
  }
  if s1.Lower() != testSHA512 {
    t.Errorf("NewSHA512Pip WriteBytes failed: expected: [%s] got [%s]", testSHA512, s1.Lower())
  }
  
  // test NewSHA512Pip.WriteString
  mp = NewSHA512Pip()
  err = mp.WriteString(testStringP1)
  if err != nil {
    t.Errorf("NewSHA512Pip WriteString failed: %v", err)
  }
  err = mp.WriteString(testStringP2)
  if err != nil {
    t.Errorf("NewSHA512Pip WriteString failed: %v", err)
  }
  s1 = mp.Result()
  if s1.Error() != nil {
    t.Errorf("NewSHA512Pip WriteString failed: %v", s1.Error())
  }
  if s1.Lower() != testSHA512 {
    t.Errorf("NewSHA512Pip WriteString failed: expected: [%s] got [%s]", testSHA512, s1.Lower())
  }
  
  // test NewSHA512Pip.Result
  mp = NewSHA512Pip()
  err = mp.WriteString(testStringP1)
  if err != nil {
    t.Errorf("NewSHA512Pip WriteString failed: %v", err)
  }
  s1 = mp.Result()
  if s1.Error() != nil {
    t.Errorf("NewSHA512Pip WriteString failed: %v", s1.Error())
  }
  if s1.Lower() != testSHA512P1 {
    t.Errorf("NewSHA512Pip WriteString failed: expected: [%s] got [%s]", testSHA512P1, s1.Lower())
  }
  err = mp.WriteString(testStringP2)
  if err != nil {
    t.Errorf("NewSHA512Pip WriteString failed: %v", err)
  }
  s1 = mp.Result()
  if s1.Error() != nil {
    t.Errorf("NewSHA512Pip WriteString failed: %v", s1.Error())
  }
  if s1.Lower() != testSHA512 {
    t.Errorf("NewSHA512Pip WriteString failed: expected: [%s] got [%s]", testSHA512, s1.Lower())
  }
}

func ExampleNewSHA512() {
  s := NewSHA512()
  res := s.FromString("Akvicor")
  if res.Error() == nil {
    fmt.Println(res.Lower())
    fmt.Println(res.Upper())
    fmt.Println(res.Slice())
  }
  
  // Output:
  // 6fdfb481af6573ac49d4031d22f0d73f0f24ea8cb113982593ec65ab085b0635b46d7683aa3f4ef36c2cb25a7bb8ba8cbae2fa3e9810d35b1c70a93f37586361
  // 6FDFB481AF6573AC49D4031D22F0D73F0F24EA8CB113982593EC65AB085B0635B46D7683AA3F4EF36C2CB25A7BB8BA8CBAE2FA3E9810D35B1C70A93F37586361
  // [111 223 180 129 175 101 115 172 73 212 3 29 34 240 215 63 15 36 234 140 177 19 152 37 147 236 101 171 8 91 6 53 180 109 118 131 170 63 78 243 108 44 178 90 123 184 186 140 186 226 250 62 152 16 211 91 28 112 169 63 55 88 99 97]
}

func ExampleNewSHA512Pip() {
  sp := NewSHA512Pip()
  io.WriteString(sp, "Akvicor")
  res := sp.Result()
  if res.Error() == nil {
    fmt.Println(res.Array())
    fmt.Println(res.Slice())
  }
  
  // Output:
  // [111 223 180 129 175 101 115 172 73 212 3 29 34 240 215 63 15 36 234 140 177 19 152 37 147 236 101 171 8 91 6 53 180 109 118 131 170 63 78 243 108 44 178 90 123 184 186 140 186 226 250 62 152 16 211 91 28 112 169 63 55 88 99 97]
  // [111 223 180 129 175 101 115 172 73 212 3 29 34 240 215 63 15 36 234 140 177 19 152 37 147 236 101 171 8 91 6 53 180 109 118 131 170 63 78 243 108 44 178 90 123 184 186 140 186 226 250 62 152 16 211 91 28 112 169 63 55 88 99 97]
}
