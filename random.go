package util

import (
  "math/rand"
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
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  res := make([]byte, length)
  for i := range res {
    res[i] = chars[r.Intn(charsLen)]
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
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  res := make([]byte, length)
  
  i := 0
  for ; i < charsLen && i < length; i++ {
    res[i] = chars[i][r.Intn(charsPerLen[i])]
  }
  for ; i < length; i++ {
    t := r.Intn(charsLen)
    res[i] = chars[t][r.Intn(charsPerLen[t])]
  }
  
  for i = range res {
    t := r.Intn(length)
    res[i], res[t] = res[t], res[i]
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
