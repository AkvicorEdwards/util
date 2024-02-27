package util

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	testString := "Akvicor"
	testStringBase64 := "QWt2aWNvcg=="

	if NewBase64().EncodeString(testString).String() != testStringBase64 {
		t.Errorf("expected %s got %s", testStringBase64, NewBase64().EncodeString(testString).String())
	}
	if NewBase64().EncodeBytes([]byte(testString)).String() != testStringBase64 {
		t.Errorf("expected %s got %s", testStringBase64, NewBase64().EncodeString(testString).String())
	}
	if !bytes.Equal(NewBase64().EncodeString(testString).Bytes(), []byte(testStringBase64)) {
		t.Errorf("expected %v got %v", []byte(testStringBase64), NewBase64().EncodeString(testString).Bytes())
	}

	if NewBase64().DecodeString(testStringBase64).Error() != nil {
		t.Errorf("decode got err: %v", NewBase64().DecodeString(testStringBase64).Error())
	}
	if NewBase64().DecodeString(testStringBase64).String() != testString {
		t.Errorf("expected %s got %s", testStringBase64, NewBase64().EncodeString(testString).String())
	}
	if NewBase64().DecodeBytes([]byte(testStringBase64)).String() != testString {
		t.Errorf("expected %s got %s", testStringBase64, NewBase64().EncodeString(testString).String())
	}
	if !bytes.Equal(NewBase64().DecodeString(testStringBase64).Bytes(), []byte(testString)) {
		t.Errorf("expected %v got %v", []byte(testStringBase64), NewBase64().EncodeString(testString).Bytes())
	}
}

func ExampleNewBase64() {
	testString := "Akvicor"
	testStringBase64 := "QWt2aWNvcg=="

	encodeS := NewBase64().EncodeString(testString)
	fmt.Println(encodeS)
	encodeB := NewBase64().EncodeBytes([]byte(testString))
	fmt.Println(encodeB.String())
	decodeS := NewBase64().DecodeString(testStringBase64)
	if decodeS.Error() == nil {
		fmt.Println(decodeS)
	}
	decodeB := NewBase64().DecodeBytes([]byte(testStringBase64))
	fmt.Println(decodeB.Bytes())

	// Output:
	// QWt2aWNvcg==
	// QWt2aWNvcg==
	// Akvicor
	// [65 107 118 105 99 111 114]
}
