package util

import (
  "crypto/sha512"
  "encoding/hex"
  "errors"
  "hash"
  "io"
  "os"
  "strings"
)

const SHA512ResultLength = 64

type SHA512Result struct {
  result [SHA512ResultLength]byte
  err    error
}

func (s *SHA512Result) Array() [SHA512ResultLength]byte {
  return s.result
}

func (s *SHA512Result) Slice() []byte {
  return s.result[:]
}

func (s *SHA512Result) Upper() string {
  return strings.ToUpper(hex.EncodeToString(s.result[:]))
}

func (s *SHA512Result) Lower() string {
  return strings.ToLower(hex.EncodeToString(s.result[:]))
}

func (s *SHA512Result) Error() error {
  return s.err
}

func NewSHA512Result(result []byte, err error) *SHA512Result {
  res := &SHA512Result{
    result: [SHA512ResultLength]byte{},
    err:    err,
  }
  copy(res.result[:], result)
  return res
}

type SHA512 struct{}

func NewSHA512() *SHA512 {
  return &SHA512{}
}

func (s *SHA512) FromReader(r io.Reader) *SHA512Result {
  ha := sha512.New()
  if _, err := io.Copy(ha, r); err != nil {
    return NewSHA512Result(nil, err)
  }
  hashInBytes := ha.Sum(nil)
  if len(hashInBytes) != SHA512ResultLength {
    return NewSHA512Result(nil, errors.New("wrong length"))
  }
  return NewSHA512Result(hashInBytes, nil)
}

func (s *SHA512) FromFile(filepath string) *SHA512Result {
  file, err := os.Open(filepath)
  if err != nil {
    return NewSHA512Result(nil, err)
  }
  defer func() {
    _ = file.Close()
  }()
  return s.FromReader(file)
}

func (s *SHA512) FromReaderChunk(r io.Reader, chunksize int) *SHA512Result {
  buf := make([]byte, chunksize)
  s1 := sha512.New()
  var n int
  var err error
  for {
    n, err = r.Read(buf)
    if err != nil {
      if err != io.EOF {
        return NewSHA512Result(nil, err)
      }
      if n > 0 {
        _, err = s1.Write(buf[:n])
        if err != nil {
          return NewSHA512Result(nil, err)
        }
      }
      break
    }
    _, err = s1.Write(buf[:n])
    if err != nil {
      return NewSHA512Result(nil, err)
    }
  }
  return NewSHA512Result(s1.Sum(nil), nil)
}

func (s *SHA512) FromFileChunk(filepath string, chunksize int) *SHA512Result {
  file, err := os.Open(filepath)
  if err != nil {
    return NewSHA512Result(nil, err)
  }
  defer func() {
    _ = file.Close()
  }()
  return s.FromReaderChunk(file, chunksize)
}

func (s *SHA512) FromBytes(b []byte) *SHA512Result {
  res := sha512.Sum512(b)
  return NewSHA512Result(res[:], nil)
}

func (s *SHA512) FromString(str string) *SHA512Result {
  res := sha512.Sum512([]byte(str))
  return NewSHA512Result(res[:], nil)
}

type SHA512Pip struct {
  sha512 hash.Hash
}

func NewSHA512Pip() *SHA512Pip {
  return &SHA512Pip{sha512: sha512.New()}
}

func (s *SHA512Pip) Write(data []byte) (n int, err error) {
  return s.sha512.Write(data)
}

func (s *SHA512Pip) WriteBytes(data []byte) error {
  _, err := s.Write(data)
  return err
}

func (s *SHA512Pip) WriteString(data string) error {
  _, err := s.Write([]byte(data))
  return err
}

func (s *SHA512Pip) Result() *SHA512Result {
  return NewSHA512Result(s.sha512.Sum(nil), nil)
}
