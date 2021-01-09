package util

import (
	"fmt"
	"log"
)

func ExampleCRC32File() {
	// dd if=/dev/zero of=testFile1 bs=3M count=2
	bs, err := CRC32File("testFile1")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(bs)

	// Output:
	// 943657036
}

func ExampleCRC32String() {
	fmt.Println(CRC32String("Hello CRC32"))

	// Output:
	// 1203018389
}

func ExampleCRC32Bytes() {
	fmt.Println(CRC32Bytes([]byte("Hello CRC32")))

	// Output:
	// 1203018389
}
