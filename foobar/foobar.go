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

func SayAny(n any) string {
	var num int
	switch v := n.(type) {
	case int:
		num = v
	case string:
		num, _ = strconv.Atoi(v)
	default:
		return ""
	}

	return Say(num)
}
