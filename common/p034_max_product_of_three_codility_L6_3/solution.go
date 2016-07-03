package main

import (
	"fmt"
	"sort"
)

const MaxUint = ^uint(0)
const maxUint = 0
const MaxInt = int(MaxUint >> 1)
const maxInt = -MaxInt - 1

func Solution(A []int) int {

	sort.Ints(A)
	// fmt.Println(A)

	// for a sorted array there might be just a couple cases were biggest triplet might be
	max := maxInt
	l := len(A)
	// therer exists 3 positive integer in the end of array
	c1 := A[l-3] * A[l-2] * A[l-1]
	if c1 > max {
		max = c1
	}
	// there exists two lowest negative multiplied on biggest positive
	c2 := A[0] * A[1] * A[l-1]
	if c2 > max {
		max = c2
	}
	// there is only two positive so that and the rest is negative
	negPos := -1
	for i, v := range A {
		if v >= 0 {
			negPos = i
			break
		}
	}

	if negPos != -1 {

		c3 := A[negPos] * A[l-2] * A[l-1]
		if c3 > max {
			max = c3
		}
	}

	// fmt.Println(max)
	return max
}

func main() {
	fmt.Println(Solution([]int{-3, 1, 2, -2, 5, 6}) == 60)
	fmt.Println(Solution([]int{10, 50, 5, 1}) == 2500)
	fmt.Println(Solution([]int{1, 1, 1}) == 1) // 0 ≤ P < Q < R < N
	fmt.Println(Solution([]int{-3, -3, -3}) == -27)
	fmt.Println(Solution([]int{2, 2, 3}) == 12)
	fmt.Println(Solution([]int{-3, 3, -3, 3, -3, 0, 3}) == 27)
	fmt.Println(Solution([]int{-3, 3, -3, -3, 0, 3}) == 27)
	fmt.Println(Solution([]int{-1, 3, 0, 3}) == 0)
	// fmt.Println(Solution([]int{-3, 3, 4, -3, 3, -3, 3, 5}) == 1)
	// fmt.Println(Solution([]int{-3, -3, -4, -3, -3, -3}) == 0)
	// fmt.Println(Solution([]int{2, 2, 2, 2, 2}) == 1)

}
