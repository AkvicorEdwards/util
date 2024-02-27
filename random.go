package util

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"math/big"
	"strings"
	"time"
)

const RandomLower = "abcdefghijklmnopqrstuvwxyz"
const RandomUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const RandomDigit = "0123456789"
const RandomSpecial = "!_.~?-+=#@$%"
const RandomAlpha = RandomLower + RandomUpper
const RandomAll = RandomAlpha + RandomDigit + RandomSpecial

var RandomSlice = []string{RandomLower, RandomUpper, RandomDigit, RandomSpecial}

func RandomString(length int, str ...string) string {
	chars := RandomAll
	if len(str) != 0 {
		chars = ""
		for _, v := range str {
			chars += v
		}
	}
	charsLen := len(chars)
	res := make([]byte, length)
	for i := range res {
		r, err := rand.Int(rand.Reader, big.NewInt(int64(charsLen)))
		if err != nil {
			return ""
		}
		res[i] = chars[r.Int64()]
	}
	return string(res)
}

func RandomStringAtLeastOnce(length int, str ...string) string {
	chars := RandomSlice
	charsLen := len(RandomSlice)
	if len(str) != 0 {
		chars = str
		charsLen = len(str)
	}
	charsPerLen := make([]int, charsLen)
	for i := range charsPerLen {
		charsPerLen[i] = len(chars[i])
	}
	res := make([]byte, length)

	i := 0
	for ; i < charsLen && i < length; i++ {
		r, err := rand.Int(rand.Reader, big.NewInt(int64(charsPerLen[i])))
		if err != nil {
			return ""
		}
		res[i] = chars[i][r.Int64()]
	}
	for ; i < length; i++ {
		r, err := rand.Int(rand.Reader, big.NewInt(int64(charsLen)))
		if err != nil {
			return ""
		}
		t := r.Int64()
		r2, err := rand.Int(rand.Reader, big.NewInt(int64(charsPerLen[t])))
		if err != nil {
			return ""
		}
		res[i] = chars[t][r2.Int64()]
	}

	for i = range res {
		t, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
		if err != nil {
			return ""
		}
		res[i], res[t.Int64()] = res[t.Int64()], res[i]
	}
	return string(res)
}

const RandomStringWithTimestampTimeLength = 8

func RandomStringWithTimestamp(length int, unix ...int64) string {
	if length < RandomStringWithTimestampTimeLength {
		length = RandomStringWithTimestampTimeLength
	}
	var t int64
	if len(unix) > 0 {
		t = unix[0]
	} else {
		t = time.Now().Unix()
	}

	return TimeUnixToBase36(t, RandomStringWithTimestampTimeLength) + RandomString(length-RandomStringWithTimestampTimeLength)
}

func ParseRandomStringWithTimestamp(str string) (int64, string) {
	if len(str) < RandomStringWithTimestampTimeLength {
		return 0, ""
	}
	t := TimeBase36ToUnix(string([]byte(str)[:RandomStringWithTimestampTimeLength]))
	if t == 0 {
		return 0, ""
	}
	return t, string([]byte(str)[RandomStringWithTimestampTimeLength:])
}

type KeyResult struct {
	key []byte
	err error
}

func (k *KeyResult) Bytes() []byte {
	if k.err != nil {
		return []byte{}
	}
	return k.key
}

func (k *KeyResult) Upper() string {
	if k.err != nil {
		return ""
	}
	return strings.ToUpper(hex.EncodeToString(k.key))
}

func (k *KeyResult) Lower() string {
	if k.err != nil {
		return ""
	}
	return strings.ToLower(hex.EncodeToString(k.key))
}

func (k *KeyResult) Error() error {
	return k.err
}

func NewKeyResult(data []byte, err error) *KeyResult {
	res := &KeyResult{
		key: make([]byte, len(data)),
		err: err,
	}
	copy(res.key, data)
	return res
}

func KeyResultFromHexString(hx string) *KeyResult {
	k, err := hex.DecodeString(hx)
	if err != nil {
		return NewKeyResult(nil, err)
	}
	return NewKeyResult(k, nil)
}

func RandomKey(l int) *KeyResult {
	k := make([]byte, l)
	n, err := rand.Read(k)
	if err != nil {
		return NewKeyResult(nil, err)
	}
	if n != l {
		return NewKeyResult(nil, errors.New("invalid length"))
	}
	return NewKeyResult(k, nil)
}
