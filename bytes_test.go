package util

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBytesCombine(t *testing.T) {
	b1 := []byte{0x11, 0x22}
	b2 := []byte{0x33, 0x44, 0x55}
	b3 := []byte{0x66}
	s := []byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66}

	b := BytesCombine(b1, b2, b3)
	if !bytes.Equal(b, s) {
		t.Errorf("expected %v got %v", s, b)
	}
}

func ExampleBytesCombine() {
	b1 := []byte{11, 22}
	b2 := []byte{33, 44, 55}
	b3 := []byte{66}

	fmt.Println(BytesCombine(b1, b2, b3))

	// Output:
	// [11 22 33 44 55 66]
}
