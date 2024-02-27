package util

import (
	"path"
	"strings"
)

func SplitPathSkip(p string, skip int) (head, tail string) {
	p = path.Clean("/" + p + "/")
	fin := 0
	cur := 0
	for skip >= 0 {
		skip--
		if fin+1 >= len(p) {
			return p, ""
		}
		cur = strings.IndexByte(p[fin+1:], '/')
		if cur < 0 {
			return p, ""
		}
		fin += cur + 1
	}
	return p[:fin], p[fin:]
}

func SplitPath(p string) (head, tail string) {
	p = path.Clean("/" + p + "/")
	return SplitPathSkip(p, 0)
}

func SplitPathRepeat(p string, repeat int) (head, tail string) {
	p = path.Clean("/" + p + "/")
	head, tail = SplitPathSkip(p, repeat)
	c := strings.Count(head, "/") - 1
	if c == repeat {
		i := strings.LastIndexByte(head, '/')
		return head[i:], tail
	} else if c < repeat {
		return "/", ""
	}
	return head, tail
}
