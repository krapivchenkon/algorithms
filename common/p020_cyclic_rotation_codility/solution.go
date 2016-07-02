package main

import (
	"fmt"
)

func ShiftSliceLeft(A []int, K int) []int {
	if len(A) == 0 {
		return A
	}
	shift := K % len(A)
	if shift == 0 {
		return A
	}
	res := A[shift:]
	for _, v := range A[:shift] {
		res = append(res, v)
	}
	return res
}

func Solution(A []int, K int) []int {
	if len(A) == 0 {
		return A
	}
	shift := K % len(A)
	if shift == 0 {
		return A
	}
	res := A[len(A)-shift:]
	for _, v := range A[:len(A)-shift] {
		res = append(res, v)
	}
	return res
}

func ShiftSliceRight(A []int, K int) []int {
	if len(A) == 0 {
		return A
	}
	shift := K % len(A)
	if shift == 0 {
		return A
	}
	res := A[len(A)-shift:]
	for _, v := range A[:len(A)-shift] {
		res = append(res, v)
	}
	return res
}

func CmpSlices(A, B []int) bool {
	if len(A) != len(B) {
		return false
	}
	return fmt.Sprintf("%v", A) == fmt.Sprintf("%v", B)
}
func main() {

	fmt.Println(CmpSlices(Solution([]int{3, 8, 9, 7, 6}, 1), []int{6, 3, 8, 9, 7}))
	fmt.Println(CmpSlices(Solution([]int{3, 8, 9, 7, 6}, 3), []int{9, 7, 6, 3, 8}))
	fmt.Println(CmpSlices(Solution([]int{3, 8, 9, 7, 6}, 13), []int{9, 7, 6, 3, 8}))

}
