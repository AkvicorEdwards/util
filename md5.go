package util

import (
  "crypto/md5"
  "encoding/hex"
  "errors"
  "hash"
  "io"
  "os"
  "strings"
)

const MD5ResultLength = 16

type MD5Result struct {
  result [MD5ResultLength]byte
  err    error
}

func (m *MD5Result) Array() [MD5ResultLength]byte {
  return m.result
}

func (m *MD5Result) Slice() []byte {
  return m.result[:]
}

func (m *MD5Result) Upper() string {
  return strings.ToUpper(hex.EncodeToString(m.result[:]))
}

func (m *MD5Result) Lower() string {
  return strings.ToLower(hex.EncodeToString(m.result[:]))
}

func (m *MD5Result) Error() error {
  return m.err
}

func NewMD5Result(result []byte, err error) *MD5Result {
  res := &MD5Result{
    result: [MD5ResultLength]byte{},
    err:    err,
  }
  copy(res.result[:], result)
  return res
}

type MD5 struct{}

func NewMD5() *MD5 {
  return &MD5{}
}

func (m *MD5) FromReader(r io.Reader) *MD5Result {
  ha := md5.New()
  if _, err := io.Copy(ha, r); err != nil {
    return NewMD5Result(nil, err)
  }
  hashInBytes := ha.Sum(nil)
  if len(hashInBytes) != MD5ResultLength {
    return NewMD5Result(nil, errors.New("wrong length"))
  }
  return NewMD5Result(hashInBytes, nil)
}

func (m *MD5) FromFile(filepath string) *MD5Result {
  file, err := os.Open(filepath)
  if err != nil {
    return NewMD5Result(nil, err)
  }
  defer func() {
    _ = file.Close()
  }()
  return m.FromReader(file)
}

func (m *MD5) FromReaderChunk(r io.Reader, chunksize int) *MD5Result {
  buf := make([]byte, chunksize)
  m5 := md5.New()
  var n int
  var err error
  for {
    n, err = r.Read(buf)
    if err != nil {
      if err != io.EOF {
        return NewMD5Result(nil, err)
      }
      if n > 0 {
        _, err = m5.Write(buf[:n])
        if err != nil {
          return NewMD5Result(nil, err)
        }
      }
      break
    }
    _, err = m5.Write(buf[:n])
    if err != nil {
      return NewMD5Result(nil, err)
    }
  }
  return NewMD5Result(m5.Sum(nil), nil)
}

func (m *MD5) FromFileChunk(filepath string, chunksize int) *MD5Result {
  file, err := os.Open(filepath)
  if err != nil {
    return NewMD5Result(nil, err)
  }
  defer func() {
    _ = file.Close()
  }()
  return m.FromReaderChunk(file, chunksize)
}

func (m *MD5) FromBytes(b []byte) *MD5Result {
  res := md5.Sum(b)
  return NewMD5Result(res[:], nil)
}

func (m *MD5) FromString(str string) *MD5Result {
  res := md5.Sum([]byte(str))
  return NewMD5Result(res[:], nil)
}

type MD5Pip struct {
  md5 hash.Hash
}

func NewMD5Pip() *MD5Pip {
  return &MD5Pip{md5: md5.New()}
}

func (m *MD5Pip) Write(data []byte) (n int, err error) {
  return m.md5.Write(data)
}

func (m *MD5Pip) WriteBytes(data []byte) error {
  _, err := m.Write(data)
  return err
}

func (m *MD5Pip) WriteString(data string) error {
  _, err := m.Write([]byte(data))
  return err
}

func (m *MD5Pip) Result() *MD5Result {
  return NewMD5Result(m.md5.Sum(nil), nil)
}
