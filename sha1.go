package util

import (
  "crypto/sha1"
  "encoding/hex"
  "errors"
  "hash"
  "io"
  "os"
  "strings"
)

const SHA1ResultLength = 20

type SHA1Result struct {
  result [SHA1ResultLength]byte
  err    error
}

func (s *SHA1Result) Array() [SHA1ResultLength]byte {
  return s.result
}

func (s *SHA1Result) Slice() []byte {
  return s.result[:]
}

func (s *SHA1Result) Upper() string {
  return strings.ToUpper(hex.EncodeToString(s.result[:]))
}

func (s *SHA1Result) Lower() string {
  return strings.ToLower(hex.EncodeToString(s.result[:]))
}

func (s *SHA1Result) Error() error {
  return s.err
}

func NewSHA1Result(result []byte, err error) *SHA1Result {
  res := &SHA1Result{
    result: [SHA1ResultLength]byte{},
    err:    err,
  }
  copy(res.result[:], result)
  return res
}

type SHA1 struct{}

func NewSHA1() *SHA1 {
  return &SHA1{}
}

func (s *SHA1) FromReader(r io.Reader) *SHA1Result {
  ha := sha1.New()
  if _, err := io.Copy(ha, r); err != nil {
    return NewSHA1Result(nil, err)
  }
  hashInBytes := ha.Sum(nil)
  if len(hashInBytes) != SHA1ResultLength {
    return NewSHA1Result(nil, errors.New("wrong length"))
  }
  return NewSHA1Result(hashInBytes, nil)
}

func (s *SHA1) FromFile(filepath string) *SHA1Result {
  file, err := os.Open(filepath)
  if err != nil {
    return NewSHA1Result(nil, err)
  }
  defer func() {
    _ = file.Close()
  }()
  return s.FromReader(file)
}

func (s *SHA1) FromReaderChunk(r io.Reader, chunksize int) *SHA1Result {
  buf := make([]byte, chunksize)
  s1 := sha1.New()
  var n int
  var err error
  for {
    n, err = r.Read(buf)
    if err != nil {
      if err != io.EOF {
        return NewSHA1Result(nil, err)
      }
      if n > 0 {
        _, err = s1.Write(buf[:n])
        if err != nil {
          return NewSHA1Result(nil, err)
        }
      }
      break
    }
    _, err = s1.Write(buf[:n])
    if err != nil {
      return NewSHA1Result(nil, err)
    }
  }
  return NewSHA1Result(s1.Sum(nil), nil)
}

func (s *SHA1) FromFileChunk(filepath string, chunksize int) *SHA1Result {
  file, err := os.Open(filepath)
  if err != nil {
    return NewSHA1Result(nil, err)
  }
  defer func() {
    _ = file.Close()
  }()
  return s.FromReaderChunk(file, chunksize)
}

func (s *SHA1) FromBytes(b []byte) *SHA1Result {
  res := sha1.Sum(b)
  return NewSHA1Result(res[:], nil)
}

func (s *SHA1) FromString(str string) *SHA1Result {
  res := sha1.Sum([]byte(str))
  return NewSHA1Result(res[:], nil)
}

type SHA1Pip struct {
  sha1 hash.Hash
}

func NewSHA1Pip() *SHA1Pip {
  return &SHA1Pip{sha1: sha1.New()}
}

func (s *SHA1Pip) Write(data []byte) (n int, err error) {
  return s.sha1.Write(data)
}

func (s *SHA1Pip) WriteBytes(data []byte) error {
  _, err := s.Write(data)
  return err
}

func (s *SHA1Pip) WriteString(data string) error {
  _, err := s.Write([]byte(data))
  return err
}

func (s *SHA1Pip) Result() *SHA1Result {
  return NewSHA1Result(s.sha1.Sum(nil), nil)
}
