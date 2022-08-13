package main

import (
	"fmt"
)

func main() {
	// "abcdef" -> []string{"ab", "cd", "ef"}
	// "abcdefg" -> []string{"ab", "cd", "ef", "g*"}
	fmt.Println(couple("abcdef"))
	fmt.Println(couple("abcdefg"))
}

func couple(s string) []string {
	var result []string

	for s += "*"; len(s) > 1; s = s[2:] {
		result = append(result, s[:2])
	}

	return result
}
