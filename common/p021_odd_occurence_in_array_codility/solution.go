package main

import (
	"fmt"
)

func Solution(A []int) int {
	if len(A) == 1 {
		return A[0]
	}
	sum := A[0]
	for i := 1; i < len(A); i++ {
		sum = sum ^ A[i]
	}
	return sum

}

func main() {

	fmt.Println(Solution([]int{9, 3, 9, 3, 9, 7, 9}) == 7)
	fmt.Println(Solution([]int{9, 3, 7, 3, 9, 7, 9}) == 9)

}
