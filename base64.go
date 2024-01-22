package util

import (
  "encoding/base64"
)

type Base64Result struct {
  result string
  err    error
}

func (b *Base64Result) Bytes() []byte {
  return []byte(b.result)
}

func (b *Base64Result) String() string {
  return b.result
}

func (b *Base64Result) Error() error {
  return b.err
}

func NewBase64Result(result string, err error) *Base64Result {
  return &Base64Result{
    result: result,
    err:    err,
  }
}

type Base64 struct{}

func NewBase64() *Base64 {
  return &Base64{}
}

func (b *Base64) EncodeBytes(data []byte) *Base64Result {
  return NewBase64Result(base64.StdEncoding.EncodeToString(data), nil)
}

func (b *Base64) EncodeString(str string) *Base64Result {
  return NewBase64Result(base64.StdEncoding.EncodeToString([]byte(str)), nil)
}

func (b *Base64) DecodeBytes(data []byte) *Base64Result {
  res, err := base64.StdEncoding.DecodeString(string(data))
  return NewBase64Result(string(res), err)
}

func (b *Base64) DecodeString(str string) *Base64Result {
  res, err := base64.StdEncoding.DecodeString(str)
  return NewBase64Result(string(res), err)
}
