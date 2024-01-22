package util

import (
  "encoding/hex"
  "hash/crc32"
  "io"
  "os"
  "strings"
)

type CRC32Result struct {
  value uint32
  err   error
}

func (c *CRC32Result) Value() uint32 {
  return c.value
}

func (c *CRC32Result) Array() [4]byte {
  return UInt32ToBytesArray(c.value)
}

func (c *CRC32Result) Slice() []byte {
  return UInt32ToBytesSlice(c.value)
}

func (c *CRC32Result) Upper() string {
  return strings.ToUpper(hex.EncodeToString(c.Slice()))
}

func (c *CRC32Result) Lower() string {
  return strings.ToLower(hex.EncodeToString(c.Slice()))
}

func (c *CRC32Result) Error() error {
  return c.err
}

func NewCRC32Result(result uint32, err error) *CRC32Result {
  return &CRC32Result{
    value: result,
    err:   err,
  }
}

type CRC32 struct{}

func NewCRC32() *CRC32 {
  return &CRC32{}
}

func (c *CRC32) FromReader(r io.Reader) *CRC32Result {
  return c.FromReaderChunk(r, 1024*8)
}

func (c *CRC32) FromFile(filepath string) *CRC32Result {
  file, err := os.Open(filepath)
  if err != nil {
    return NewCRC32Result(0, err)
  }
  defer func() {
    _ = file.Close()
  }()
  return c.FromReader(file)
}

func (c *CRC32) FromReaderChunk(r io.Reader, chunksize int) *CRC32Result {
  val := uint32(0)
  buf := make([]byte, chunksize)
  var n int
  var err error
  for {
    n, err = r.Read(buf)
    if err != nil {
      if err != io.EOF {
        return NewCRC32Result(0, err)
      }
      if n > 0 {
        val = crc32.Update(val, crc32.IEEETable, buf[:n])
      }
      break
    }
    val = crc32.Update(val, crc32.IEEETable, buf[:n])
  }
  return NewCRC32Result(val, nil)
}

func (c *CRC32) FromFileChunk(filepath string, chunksize int) *CRC32Result {
  file, err := os.Open(filepath)
  if err != nil {
    return NewCRC32Result(0, err)
  }
  defer func() {
    _ = file.Close()
  }()
  return c.FromReaderChunk(file, chunksize)
}

func (c *CRC32) FromBytes(data []byte) *CRC32Result {
  return NewCRC32Result(crc32.ChecksumIEEE(data), nil)
}

func (c *CRC32) FromString(data string) *CRC32Result {
  return NewCRC32Result(crc32.ChecksumIEEE([]byte(data)), nil)
}

type CRC32Pip struct {
  crc uint32
}

func NewCRC32Pip() *CRC32Pip {
  return &CRC32Pip{crc: 0}
}

func (c *CRC32Pip) Write(data []byte) (n int, err error) {
  c.crc = crc32.Update(c.crc, crc32.IEEETable, data)
  return len(data), nil
}

func (c *CRC32Pip) WriteBytes(data []byte) {
  c.crc = crc32.Update(c.crc, crc32.IEEETable, data)
}

func (c *CRC32Pip) WriteString(data string) {
  c.crc = crc32.Update(c.crc, crc32.IEEETable, []byte(data))
}

func (c *CRC32Pip) Result() *CRC32Result {
  return NewCRC32Result(c.crc, nil)
}
