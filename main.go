package main

import (
	"fmt"
	"os"
)

var name string = "Builder"

func main() {
	n := os.Getenv("NAME")
	if n != "" {
		name = n
	}
	fmt.Println("Hello,", name)
}
