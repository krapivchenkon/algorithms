package main

import (
	"fmt"
	"sort"
)

func Solution(A []int) int {

	A1 := make([]int, len(A))
	A2 := make([]int, len(A))

	for i, v := range A {
		A1[i] = v + i
		A2[i] = i - v
	}
	sort.Ints(A1)
	sort.Ints(A2)
	fmt.Println(A1, A2)
	// fmt.Println(A2)
	inter := 0
	pos := 0
	for _, v := range A1 {
		pos = sort.SearchInts(A2, v)
		if pos == 0 {
			inter += 1
		} else if pos == len(A2) {
			// element not found
			inter += pos

		} else if v == A2[pos] {
			// element exists in array
			inter += pos + 1
		} else {
			//element not found but pos points
			// to the first element that is larger than v
			inter += pos
		}

	}

	inter = inter - len(A)*(len(A)+1)/2

	if inter > 10000000 {
		inter = -1
	}
	return inter
}

func main() {
	fmt.Println(Solution([]int{1, 5, 2, 1, 4, 0}) == 11)
	fmt.Println(Solution([]int{1, 0, 1, 0, 1}) == 6)
	// fmt.Println(Solution([]int{1, 1, 1}) == 1) // 0 ≤ P < Q < R < N
	// fmt.Println(Solution([]int{-3, -3, -3}) == -27)
	// fmt.Println(Solution([]int{2, 2, 3}) == 12)
	// fmt.Println(Solution([]int{-3, 3, -3, 3, -3, 0, 3}) == 27)
	// fmt.Println(Solution([]int{-3, 3, -3, -3, 0, 3}) == 27)
	// fmt.Println(Solution([]int{-1, 3, 0, 3}) == 0)
	// fmt.Println(Solution([]int{-3, 3, 4, -3, 3, -3, 3, 5}) == 1)
	// fmt.Println(Solution([]int{-3, -3, -4, -3, -3, -3}) == 0)
	// fmt.Println(Solution([]int{2, 2, 2, 2, 2}) == 1)

}
