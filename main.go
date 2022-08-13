package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"G": 71,
		"O": 79,
		"P": 80,
		"H": 72,
		"E": 69,
		"R": 82,
	}
	var keys = []string{}
	var vals = []int{}

	for i, v := range m {
		keys = append(keys, i)
		vals = append(vals, v)
	}

	fmt.Println(keys)
	fmt.Println(vals)

}
