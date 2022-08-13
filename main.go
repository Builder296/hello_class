package main

import "fmt"

func main() {
	matrix := [4][4]int{
		{1, 3, 4, 2},
		{3, 4, 2, 1},
		{4, 2, 1, 3},
		{2, 1, 3, 4},
	}

	r := multiply(matrix, 2)

	for _, rr := range r {
		fmt.Println(rr)
	}
}

func multiply(matix [4][4]int, n int) [4][4]int {
	var result [4][4]int
	for i, nSet := range matix {
		var resultN [4]int
		for j, num := range nSet {
			resultN[j] = num * n
		}
		result[i] = resultN
	}
	return result
}
