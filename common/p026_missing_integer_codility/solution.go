package main

import (
	"fmt"
)

func Solution(A []int) int {
	// minInt := 1
	if len(A) == 1 {
		if A[0] == 1 {
			return 2
		} else {
			return 1
		}
	}
	ind := make(map[int]struct{})

	for _, v := range A {
		// fill map with unique indeces
		if _, ok := ind[v]; !ok && v > 0 {
			ind[v] = struct{}{}
		}
	}

	for i := 1; i <= len(A)+1; i++ {
		// iterate over key incrementing them
		if _, ok := ind[i]; !ok {

			return i
		}
	}
	return 1
}

func main() {

	fmt.Println(Solution([]int{1, 3, 6, 4, 1, 2}) == 5)
	fmt.Println(Solution([]int{-4, -5, -1000, 2}) == 1)
	fmt.Println(Solution([]int{-4, -5, -1000, 1}) == 2)
	fmt.Println(Solution([]int{1, 2, 3, 4, 5, 6}) == 7)

	fmt.Println(Solution([]int{100}) == 1)
	fmt.Println(Solution([]int{1}) == 2)

}
