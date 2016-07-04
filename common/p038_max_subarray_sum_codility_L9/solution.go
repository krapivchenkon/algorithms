package main

import (
	"fmt"
)

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Implementation of Kadane's algorithm

func Solution(A []int) int {
	if len(A) == 1 {
		// one element in the array
		return A[0]
	}

	maxCur, max := A[0], A[0]

	for i := 1; i < len(A); i++ {
		maxCur = Max(A[i], maxCur+A[i])
		max = Max(maxCur, max)
	}
	// fmt.Println(max)
	return max
}

func GetMaxSlice(A []int) []int {
	// this solution actually returns slice with max consequitive sum
	// modification of Kadane's algorithm
	if len(A) == 1 {
		// one element in the array
		return A
	}
	maxCur, max := A[0], A[0]
	st, end, stCur, endCur := 0, 0, 0, 0

	for i := 1; i < len(A); i++ {
		// if maxCur+A[i] < 0 {
		// 	stCur, endCur = i, i
		// }
		if A[i] < maxCur+A[i] {
			maxCur = maxCur + A[i]
			endCur = i
		} else {
			maxCur = A[i]
			stCur, endCur = i, i

		}
		if max < maxCur {
			st, end = stCur, endCur
			// st = stCur
			max = maxCur
		}

	}
	// fmt.Println(st, end, max, A[st:end+1])
	return A[st : end+1]
}

func CmpSlices(A, B []int) bool {
	if len(A) != len(B) {
		return false
	}
	return fmt.Sprintf("%v", A) == fmt.Sprintf("%v", B)
}

func main() {
	fmt.Println(Solution([]int{3, -1, 5, -6, -9}) == 7)
	fmt.Println(Solution([]int{-3, -1, -5, -6, -9}) == -1) // all negative
	fmt.Println(Solution([]int{1, 0, 1, 0, 1}) == 3)       //

	fmt.Println(Solution([]int{4, 3, 4, 4, 4, 2}) == 21) // all positive

	fmt.Println(Solution([]int{4}) == 4)   // one element
	fmt.Println(Solution([]int{-3}) == -3) // one negative element

	fmt.Println(Solution([]int{-100, 100, -1, 100, -100}) == 199)
	fmt.Println(Solution([]int{-2, 3, 2, -1}) == 5)
	fmt.Println(Solution([]int{1, -3, 2, 1, -1}) == 3)
	fmt.Println(Solution([]int{-2, -3, 4, -1, -2, 1, 5, -3}) == 7)

	fmt.Println("get indeces:")
	fmt.Println(CmpSlices(GetMaxSlice([]int{3, -1, 5, -6, -9}), []int{3, -1, 5}))
	fmt.Println(CmpSlices(GetMaxSlice([]int{-3, -1, -5, -6, -9}), []int{-1}))
	fmt.Println(CmpSlices(GetMaxSlice([]int{1, 0, 1, 0, 1}), []int{1, 0, 1, 0, 1}))
	fmt.Println(CmpSlices(GetMaxSlice([]int{4, 3, 4, 4, 4, 2}), []int{4, 3, 4, 4, 4, 2}))
	fmt.Println(CmpSlices(GetMaxSlice([]int{4}), []int{4}))
	fmt.Println(CmpSlices(GetMaxSlice([]int{-3}), []int{-3}))
	fmt.Println(CmpSlices(GetMaxSlice([]int{-100, 100, -1, 100, -100}), []int{100, -1, 100}))
	fmt.Println(CmpSlices(GetMaxSlice([]int{-2, 3, 2, -1}), []int{3, 2}))
	fmt.Println(CmpSlices(GetMaxSlice([]int{1, -3, 2, 1, -1}), []int{2, 1}))
	fmt.Println(CmpSlices(GetMaxSlice([]int{-2, -3, 4, -1, -2, 1, 5, -3}), []int{4, -1, -2, 1, 5}))
	fmt.Println(CmpSlices(GetMaxSlice([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}), []int{4, -1, 2, 1}))

}
