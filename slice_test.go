package util

import (
	"fmt"
)

func ExampleRemoveDuplicates() {
	s := []string{
		"aaa",
		"bbb",
		"aaa",
		"ccc",
		"aaa",
	}
	ss := NewRemoveDuplicates().String(s)
	fmt.Println(ss)

	// Output:
	// [aaa bbb ccc]
}
