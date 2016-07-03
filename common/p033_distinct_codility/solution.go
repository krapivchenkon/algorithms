package main

import (
	"fmt"
	// "sort"
)

func Solution(A []int) int {

	ind := make(map[int]struct{})

	for _, v := range A {
		if _, ok := ind[v]; !ok {
			// key doesn't exists
			ind[v] = struct{}{}
		}
	}

	return len(ind)
}

func main() {
	fmt.Println(Solution([]int{2, 1, 1, 2, 3, 1}) == 3)
	fmt.Println(Solution([]int{1, 1, 1, 1, 1}) == 1)
	fmt.Println(Solution([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) == 10)
	fmt.Println(Solution([]int{}) == 0)
	fmt.Println(Solution([]int{-3, 3, -3, 3, -3, 3}) == 2)
	fmt.Println(Solution([]int{-3, 3, 4, -3, 3, -3, 3, 5}) == 4)
	fmt.Println(Solution([]int{2, 2, 2, 2, 2}) == 1)

}
