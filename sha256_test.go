package util

import (
	"fmt"
	"log"
)

func ExampleSHA256File() {
	// dd if=/dev/zero of=testFile1 bs=3M count=2
	bs, err := SHA256File("testFile1")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%x\n", bs)

	// Output:
	// b69dae56a14d1a8314ed40664c4033ea0a550eea2673e04df42a66ac6b9faf2c
}

func ExampleSHA256String() {
	fmt.Printf("%x\n", SHA256String("Hello Sha256"))

	// Output:
	// b689b7eb84fe084725de196e09481d4fad8d2c50c99fb6a33739543a5ca250e2
}

func ExampleSHA256Bytes() {
	fmt.Printf("%x\n", SHA256Bytes([]byte("Hello Sha256")))

	// Output:
	// b689b7eb84fe084725de196e09481d4fad8d2c50c99fb6a33739543a5ca250e2
}
