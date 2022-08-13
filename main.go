package main

import "fmt"

func main() {
	fmt.Println("isPalindrome:", isPalindrome(1, 2, 3, 3, 1))
}

func isPalindrome(a, b, c, d, e int) bool {
	return a == e && b == d
}
