package util

import "fmt"

func ExampleBytesCombine() {
	bs1 := []byte{11, 22}
	bs2 := []byte{33, 44}
	bs3 := BytesCombine(bs1, bs2)
	fmt.Println(bs3)

	// Output:
	// [11 22 33 44]
}

func ExampleBitSet() {
	b8 := uint8(0)
	b16 := uint16(0)
	b32 := uint32(0)
	b64 := uint64(0)
	b := uint(0)
	BitSet(&b8, 0, true)
	BitSet(&b16, 0, true)
	BitSet(&b32, 0, true)
	BitSet(&b64, 0, true)
	BitSet(&b, 0, true)
	fmt.Println(b8)
	fmt.Println(b16)
	fmt.Println(b32)
	fmt.Println(b64)
	fmt.Println(b)

	BitSet(&b8, 7, true)
	BitSet(&b16, 15, true)
	BitSet(&b32, 31, true)
	BitSet(&b64, 63, true)
	BitSet(&b, 1, true)
	fmt.Println(b8)
	fmt.Println(b16)
	fmt.Println(b32)
	fmt.Println(b64)
	fmt.Println(b)

	BitSet(&b8, 0, false)
	BitSet(&b16, 0, false)
	BitSet(&b32, 0, false)
	BitSet(&b64, 0, false)
	BitSet(&b, 0, false)
	fmt.Println(b8)
	fmt.Println(b16)
	fmt.Println(b32)
	fmt.Println(b64)
	fmt.Println(b)

	e := "e"
	BitSet(&e, 0, true)
	fmt.Println(e)

	// Output:
	// 1
	// 1
	// 1
	// 1
	// 1
	// 129
	// 32769
	// 2147483649
	// 9223372036854775809
	// 3
	// 128
	// 32768
	// 2147483648
	// 9223372036854775808
	// 2
	// e
}

func ExampleSplitPath() {
	test := func(s string) {
		head, tail := SplitPath(s)
		fmt.Printf("s[%s] -> [%s, %s]\n", s, head, tail)
	}
	test("/123/akvicor")
	test("123/akvicor")
	test("/akvicor")
	test("akvicor")
	test("/")
	test("")

	// Output:
	// s[/123/akvicor] -> [/123, /akvicor]
	// s[123/akvicor] -> [/123, /akvicor]
	// s[/akvicor] -> [/akvicor, ]
	// s[akvicor] -> [/akvicor, ]
	// s[/] -> [/, ]
	// s[] -> [/, ]
}

func ExampleSplitPathSkip() {
	test := func(s string, skip int) {
		head, tail := SplitPathSkip(s, skip)
		fmt.Printf("s[%s] skip[%d] -> [%s, %s]\n", s, skip, head, tail)
	}
	s := "/1/23/456/7/"
	test(s, 0)
	test(s, 1)
	test(s, 2)
	test(s, 3)
	test(s, 4)
	s = ""
	test(s, 0)
	test(s, 1)
	s = "/"
	test(s, 0)
	test(s, 1)
	s = "url"
	test(s, 0)
	test(s, 1)
	s = "/url"
	test(s, 0)
	test(s, 1)
	s = "url/"
	test(s, 0)
	test(s, 1)
	s = "/url/"
	test(s, 0)
	test(s, 1)

	// Output:
	// s[/1/23/456/7/] skip[0] -> [/1, /23/456/7]
	// s[/1/23/456/7/] skip[1] -> [/1/23, /456/7]
	// s[/1/23/456/7/] skip[2] -> [/1/23/456, /7]
	// s[/1/23/456/7/] skip[3] -> [/1/23/456/7, ]
	// s[/1/23/456/7/] skip[4] -> [/1/23/456/7, ]
	// s[] skip[0] -> [/, ]
	// s[] skip[1] -> [/, ]
	// s[/] skip[0] -> [/, ]
	// s[/] skip[1] -> [/, ]
	// s[url] skip[0] -> [/url, ]
	// s[url] skip[1] -> [/url, ]
	// s[/url] skip[0] -> [/url, ]
	// s[/url] skip[1] -> [/url, ]
	// s[url/] skip[0] -> [/url, ]
	// s[url/] skip[1] -> [/url, ]
	// s[/url/] skip[0] -> [/url, ]
	// s[/url/] skip[1] -> [/url, ]
}

func ExampleRandomString() {
	fmt.Printf("[%s]\n", RandomString(0, "a"))
	fmt.Printf("[%s]\n", RandomString(0, "abc"))
	fmt.Printf("[%s]\n", RandomString(1, "b"))
	fmt.Printf("[%s]\n", RandomString(7, "a"))
	fmt.Printf("[%s]\n", RandomString(0, "a", "b"))

	// Output:
	// []
	// []
	// [b]
	// [aaaaaaa]
	// [bbbbbbbb]
	// []
}

func ExampleRandomStringWithTimestamp() {
	rstr := RandomStringWithTimestamp(17)
	fmt.Printf("[%s]\n", rstr)
	date, str := ParseRandomStringWithTimestamp(rstr)
	fmt.Printf("[%d][%s]\n", date, str)
}
