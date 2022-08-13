package main

import (
	"fmt"
)

func main() {
	// "abcdef" -> []string{"ab", "cd", "ef"}
	// "abcdefg" -> []string{"ab", "cd", "ef", "g*"}
	fmt.Println(couple("abcdef"))
	fmt.Println(couple("abcdefg"))
	fmt.Println(couple("สวัสดีครับ"))
}

func couple(str string) []string {
	var result []string
	s := []rune(str)
	for s = append(s, []rune("*")...); len(s) > 1; s = s[2:] {
		result = append(result, string(s[:2]))
	}

	return result
}
