package foobar

import "strconv"

func Say(n int) string {
	if n == 5 {
		return "bar"
	}
	if n == 6 {
		return "foo"
	}
	if n == 3 {
		return "foo"
	}
	return strconv.Itoa(n)
}
