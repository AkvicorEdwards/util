package util

import (
  "fmt"
  "time"
)

func ExampleRandomString() {
  if RandomString(0, "a") != "" {
    fmt.Println("1 failed")
  }
  if RandomString(0, "abc") != "" {
    fmt.Println("2 failed")
  }
  if RandomString(1, "b") != "b" {
    fmt.Println("3 failed")
  }
  if RandomString(7, "a") != "aaaaaaa" {
    fmt.Println("4 failed")
  }
  if RandomString(0, "a", "b") != "" {
    fmt.Println("5 failed")
  }
  str := RandomString(7)
  if len(str) != 7 {
    fmt.Println("6 failed")
  }
  
  // Output:
  //
}

func ExampleRandomStringAtLeastOnce() {
  if RandomStringAtLeastOnce(0, "a") != "" {
    fmt.Println("1 failed")
  }
  if RandomStringAtLeastOnce(0, "abc") != "" {
    fmt.Println("2 failed")
  }
  if RandomStringAtLeastOnce(1, "b") != "b" {
    fmt.Println("3 failed")
  }
  if RandomStringAtLeastOnce(7, "a") != "aaaaaaa" {
    fmt.Println("4 failed")
  }
  if RandomStringAtLeastOnce(0, "a", "b") != "" {
    fmt.Println("5 failed")
  }
  str := RandomStringAtLeastOnce(2, "a", "b")
  if str != "ab" && str != "ba" {
    fmt.Println("6 failed", str)
  }
  if RandomStringAtLeastOnce(1, "a", "b") != "a" {
    fmt.Println("7 failed")
  }
  str = RandomStringAtLeastOnce(7)
  if len(str) != 7 {
    fmt.Println("8 failed")
  }
  
  // Output:
  //
}

func ExampleRandomStringWithTimestamp() {
  t := time.Now().Unix()
  rstr := RandomStringWithTimestamp(17, t)
  date, str := ParseRandomStringWithTimestamp(rstr)
  if t != date || str != rstr[RandomStringWithTimestampTimeLength:] {
    fmt.Printf("before [%d] [%s] after [%d][%s]\n", t, rstr, date, str)
  }
  
  // Output:
  //
}
