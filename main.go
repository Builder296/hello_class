package main

import "fmt"

func main() {
	fmt.Println("power:", power(2, 5))
}

func power(b, x int) int {
	result := 1
	for i := 0; i < x; i++ {
		result *= b
	}
	return result
}

func mainBasic() {
	println(squareArea(20))
	fmt.Println("Aquare area of 3 is", squareArea(20))

	a, b := swap(1, 5)
	fmt.Println("Swap 1 and 2:", a, b)

	if ok := IsCorrect(); ok {
		println("It's correct")
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
