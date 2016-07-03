package main

import (
	"fmt"
	"sort"
)

func Solution(A []int) int {
	// res := 0

	sort.Ints(A)
	// fmt.Println(A)
	for i := len(A) - 1; i >= 2; i-- {
		if A[i] <= 0 || A[i-1] <= 0 || A[i-2] <= 0 {
			// one of the values is negative
			return 0
		}
		// triangle constraints are:
		// A[P] + A[Q] > A[R],
		// A[Q] + A[R] > A[P],
		// A[R] + A[P] > A[Q].
		// but for sorted array we need to make
		// sure only first constraint is valid
		if (A[i-1] + A[i-2]) > A[i] {
			return 1
		}

	}
	// fmt.Println(res)
	return 0
}

func main() {
	fmt.Println(Solution([]int{10, 2, 5, 1, 8, 20}) == 1)
	fmt.Println(Solution([]int{10, 50, 5, 1}) == 0)
	fmt.Println(Solution([]int{1, 2, 6}) == 0)
	fmt.Println(Solution([]int{4, 2, 6}) == 0)
	fmt.Println(Solution([]int{4, 3, 6}) == 1)
	fmt.Println(Solution([]int{-3, 3, -3, 3, -3, 3}) == 1)
	fmt.Println(Solution([]int{-3, 3, 4, -3, 3, -3, 3, 5}) == 1)
	fmt.Println(Solution([]int{-3, -3, -4, -3, -3, -3}) == 0)
	fmt.Println(Solution([]int{2, 2, 2, 2, 2}) == 1)

}
