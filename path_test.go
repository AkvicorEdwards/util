package util

import "fmt"

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

func ExampleSplitPathRepeat() {
	test := func(s string, repeat int) {
		head, tail := SplitPathRepeat(s, repeat)
		fmt.Printf("s[%s] repeat[%d] -> [%s, %s]\n", s, repeat, head, tail)
	}
	s := "/1/23/456/7/"
	test(s, 0)
	test(s, 1)
	test(s, 2)
	test(s, 3)
	test(s, 4)
	test(s, 5)
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
	// s[/1/23/456/7/] repeat[0] -> [/1, /23/456/7]
	// s[/1/23/456/7/] repeat[1] -> [/23, /456/7]
	// s[/1/23/456/7/] repeat[2] -> [/456, /7]
	// s[/1/23/456/7/] repeat[3] -> [/7, ]
	// s[/1/23/456/7/] repeat[4] -> [/, ]
	// s[/1/23/456/7/] repeat[5] -> [/, ]
	// s[] repeat[0] -> [/, ]
	// s[] repeat[1] -> [/, ]
	// s[/] repeat[0] -> [/, ]
	// s[/] repeat[1] -> [/, ]
	// s[url] repeat[0] -> [/url, ]
	// s[url] repeat[1] -> [/, ]
	// s[/url] repeat[0] -> [/url, ]
	// s[/url] repeat[1] -> [/, ]
	// s[url/] repeat[0] -> [/url, ]
	// s[url/] repeat[1] -> [/, ]
	// s[/url/] repeat[0] -> [/url, ]
	// s[/url/] repeat[1] -> [/, ]
}
