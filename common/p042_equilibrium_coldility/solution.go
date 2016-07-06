package main

import (
	"fmt"
)

func Solution(A []int) int {

	if len(A) == 0 {
		// no equilibrium for empty arary
		return -1
	}

	if len(A) == 1 {
		// no equilibrium for one element arrary
		return 0
	}

	pref := make([]int, len(A)+1)
	suf := make([]int, len(A)+1)

	for i := 1; i < len(pref); i++ {
		pref[i] = pref[i-1] + A[i-1]
		suf[len(suf)-1-i] = suf[len(suf)-i] + A[len(A)-i]

	}
	// fmt.Println(pref)
	// fmt.Println(suf)

	for j := 0; j < len(pref)-1; j++ {
		if pref[j] == suf[j+1] {
			return j
			// fmt.Println(j)
		}
	}

	return -1
}

func main() {

	fmt.Println(Solution([]int{-1, 3, -4, 5, 1, -6, 2, 1}) == 1)
	fmt.Println(Solution([]int{0, 1, 1, 0, 1}) == 2)

	fmt.Println(Solution([]int{}) == -1)
	fmt.Println(Solution([]int{10}) == 0)
	fmt.Println(Solution([]int{1, 1}) == -1)

	fmt.Println(Solution([]int{3, 1, 2, 4, 3}) == -1) // no equi
	fmt.Println(Solution([]int{3, 1, 2, 4, 3, 1, 1, 1}) == 3)
	fmt.Println(Solution([]int{3, 5}) == -1)
	fmt.Println(Solution([]int{3, -3}) == -1)
	fmt.Println(Solution([]int{3, -3, 1}) == 2)
	fmt.Println(Solution([]int{3, 0}) == 0)

}
