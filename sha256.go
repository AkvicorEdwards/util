package util

import (
  "crypto/sha256"
  "encoding/hex"
  "errors"
  "hash"
  "io"
  "os"
  "strings"
)

const SHA256ResultLength = 32

type SHA256Result struct {
  result [SHA256ResultLength]byte
  err    error
}

func (s *SHA256Result) Array() [SHA256ResultLength]byte {
  return s.result
}

func (s *SHA256Result) Slice() []byte {
  return s.result[:]
}

func (s *SHA256Result) Upper() string {
  return strings.ToUpper(hex.EncodeToString(s.result[:]))
}

func (s *SHA256Result) Lower() string {
  return strings.ToLower(hex.EncodeToString(s.result[:]))
}

func (s *SHA256Result) Error() error {
  return s.err
}

func NewSHA256Result(result []byte, err error) *SHA256Result {
  res := &SHA256Result{
    result: [SHA256ResultLength]byte{},
    err:    err,
  }
  copy(res.result[:], result)
  return res
}

type SHA256 struct{}

func NewSHA256() *SHA256 {
  return &SHA256{}
}

func (s *SHA256) FromReader(r io.Reader) *SHA256Result {
  ha := sha256.New()
  if _, err := io.Copy(ha, r); err != nil {
    return NewSHA256Result(nil, err)
  }
  hashInBytes := ha.Sum(nil)
  if len(hashInBytes) != SHA256ResultLength {
    return NewSHA256Result(nil, errors.New("wrong length"))
  }
  return NewSHA256Result(hashInBytes, nil)
}

func (s *SHA256) FromFile(filepath string) *SHA256Result {
  file, err := os.Open(filepath)
  if err != nil {
    return NewSHA256Result(nil, err)
  }
  defer func() {
    _ = file.Close()
  }()
  return s.FromReader(file)
}

func (s *SHA256) FromReaderChunk(r io.Reader, chunksize int) *SHA256Result {
  buf := make([]byte, chunksize)
  s1 := sha256.New()
  var n int
  var err error
  for {
    n, err = r.Read(buf)
    if err != nil {
      if err != io.EOF {
        return NewSHA256Result(nil, err)
      }
      if n > 0 {
        _, err = s1.Write(buf[:n])
        if err != nil {
          return NewSHA256Result(nil, err)
        }
      }
      break
    }
    _, err = s1.Write(buf[:n])
    if err != nil {
      return NewSHA256Result(nil, err)
    }
  }
  return NewSHA256Result(s1.Sum(nil), nil)
}

func (s *SHA256) FromFileChunk(filepath string, chunksize int) *SHA256Result {
  file, err := os.Open(filepath)
  if err != nil {
    return NewSHA256Result(nil, err)
  }
  defer func() {
    _ = file.Close()
  }()
  return s.FromReaderChunk(file, chunksize)
}

func (s *SHA256) FromBytes(b []byte) *SHA256Result {
  res := sha256.Sum256(b)
  return NewSHA256Result(res[:], nil)
}

func (s *SHA256) FromString(str string) *SHA256Result {
  res := sha256.Sum256([]byte(str))
  return NewSHA256Result(res[:], nil)
}

type SHA256Pip struct {
  sha256 hash.Hash
}

func NewSHA256Pip() *SHA256Pip {
  return &SHA256Pip{sha256: sha256.New()}
}

func (s *SHA256Pip) Write(data []byte) (n int, err error) {
  return s.sha256.Write(data)
}

func (s *SHA256Pip) WriteBytes(data []byte) error {
  _, err := s.Write(data)
  return err
}

func (s *SHA256Pip) WriteString(data string) error {
  _, err := s.Write([]byte(data))
  return err
}

func (s *SHA256Pip) Result() *SHA256Result {
  return NewSHA256Result(s.sha256.Sum(nil), nil)
}
