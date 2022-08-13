package main

import "fmt"

func main() {
	var a []string
	fmt.Println(a == nil)

	a = make([]string, 1, 2)
	fmt.Println(a)

	a[0] = "Hello"
	fmt.Println(a)
}
