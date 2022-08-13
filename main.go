package main

import "fmt"

func main() {
	var a []string
	fmt.Println(a == nil)

	a[0] = "Hello"
	fmt.Println(a)
}
