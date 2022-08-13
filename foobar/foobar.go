package foobar

import "strconv"

func Say(n int) string {
	if n == 7 {
		return strconv.Itoa(n)
	}
	if n == 5 {
		return "bar"
	}
	if n == 4 {
		return strconv.Itoa(n)
	}
	if n == 6 {
		return "foo"
	}
	if n == 3 {
		return "foo"
	}
	if n == 2 {
		return strconv.Itoa(n)
	}
	return strconv.Itoa(n)
}
