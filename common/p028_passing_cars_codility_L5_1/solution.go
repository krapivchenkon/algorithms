package main

import (
	"fmt"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func Solution(A []int) int {
	if len(A) == 1 {
		// zero pair in case of one element array
		return 0
	}
	// get sum of 1 in A
	ones := 0
	for _, v := range A {
		if v == 1 {
			ones += 1
		}
	}
	pairs := 0
	// iterate over A once again calculating pairs and decreasing ones counter
	for _, v := range A {
		if v == 1 {
			ones -= 1
		} else {
			pairs += ones
			if pairs > 1000000000 {
				return -1
			}
		}

	}

	return pairs

}

func main() {

	fmt.Println(Solution([]int{0, 1, 0, 1, 1}) == 5)
	fmt.Println(Solution([]int{0}) == 0)
	fmt.Println(Solution([]int{1}) == 0)
	fmt.Println(Solution([]int{0, 0, 0, 0, 0, 0, 0}) == 0)
	fmt.Println(Solution([]int{1, 1, 1, 1, 1, 1}) == 0)

}
