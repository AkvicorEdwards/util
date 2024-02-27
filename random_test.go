package util

import (
	"bytes"
	"fmt"
	"testing"
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

func TestKey(t *testing.T) {
	key := RandomKey(17)
	if key.Error() != nil {
		t.Error(key.Error())
	}
	if len(key.Bytes()) != 17 {
		t.Errorf("expected 17 got %d", len(key.Bytes()))
	}
	if len(key.Upper()) != 34 {
		t.Errorf("expected 34 got %d", len(key.Upper()))
	}

	dk := KeyResultFromHexString(key.Upper())
	if dk.Error() != nil {
		t.Error(dk.Error())
	}
	if len(dk.Upper()) != 34 {
		t.Errorf("expected 34 got %d", len(dk.Upper()))
	}

	if dk.Upper() != key.Upper() {
		t.Error("low: dk != key")
	}
	if dk.Lower() != key.Lower() {
		t.Error("up: dk != key")
	}
	if !bytes.Equal(dk.Bytes(), key.Bytes()) {
		t.Error("bytes: dk != key")
	}
}
