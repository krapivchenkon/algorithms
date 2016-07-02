package main

import (
	"fmt"
)

func Solution(A []int) int {
	// if A is empty return 1
	// sum of all numbers between 1..N+1
	total := (len(A) + 2) * (len(A) + 1) / 2
	actual := 0
	for _, v := range A {
		actual += v
	}

	return total - actual
}

func main() {
	fmt.Println(Solution([]int{2, 3, 1, 5}) == 4)

	fmt.Println(Solution([]int{}) == 1)
}
