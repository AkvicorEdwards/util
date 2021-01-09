package util

import (
	"fmt"
	"log"
)

func ExampleMD5File() {
	// dd if=/dev/zero of=testFile1 bs=3M count=2
	bs, err := MD5File("testFile1")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%x\n", bs)

	// Output:
	// da6a0d097e307ac52ed9b4ad551801fc
}

func ExampleMD5String() {
	fmt.Printf("%x\n", MD5String("Hello MD5"))

	// Output:
	// e5dadf6524624f79c3127e247f04b548
}

func ExampleMD5Bytes() {
	fmt.Printf("%x\n", MD5Bytes([]byte("Hello MD5")))

	// Output:
	// e5dadf6524624f79c3127e247f04b548
}
