package main

import (
	"fmt"
	"strings"
)

type String string

func (s String) toUpper() string {
	return strings.ToUpper(string(s))
}

func main() {

	var a String = "abc"

	fmt.Println(a.toUpper())
}
