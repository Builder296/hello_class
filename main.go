package main

import "fmt"

func main() {
	println(squareArea(20))
	fmt.Println("Aquare area of 3 is", squareArea(20))

	a, b := swap(1, 5)
	fmt.Println("Swap 1 and 2:", a, b)

	if ok := IsCorrect(); ok { // scope
		println("It's correct")
	}
	{
		a := 10
		println("a:", a)
	}
}

func squareArea(a float64) float64 {
	return a * a
}

func swap(a, b int) (int, int) {
	return b, a
}

func IsCorrect() bool {
	return true
}
