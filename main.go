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

	// Go https://go.dev/tour/moretypes/18

}

func couple(str string) []string {
	var r []string
	s := []rune(str)
	for s = append(s, []rune("*")...); len(s) > 1; s = s[2:] {
		r = append(r, string(s[:2]))
	}

	return r
}
