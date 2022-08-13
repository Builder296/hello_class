package main

import (
	"fmt"
	"strings"
)

type String string

func (s *String) toUpper() {
	*s = String(strings.ToUpper(string(*s)))
}

func main() {

	var a String = "abc"
	a.toUpper()
	fmt.Println(a)
}
