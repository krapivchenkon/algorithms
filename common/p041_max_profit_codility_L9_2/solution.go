package main

import (
	"fmt"
	// "math"
)

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Solution(A []int) int {

	max, cur := 0, 0
	p, q := 0, 0
	for i := 0; i < len(A); i++ {
		cur = A[q] - A[p]
		// fmt.Println(cur, A[q], A[p])
		if cur < 0 {
			p, q = i, i+1
		} else {
			q++
		}
		max = Max(max, cur)
	}
	// fmt.Println(p, q)
	if max < 0 {
		return 0
	}
	return max
}

func main() {
	fmt.Println(Solution([]int{23171, 21011, 21123, 21366, 21013, 21367}) == 356)
	fmt.Println(Solution([]int{}) == 0)
	fmt.Println(Solution([]int{3, 2, 6, 1, 4, 5, 1, 2}) == 4)
	fmt.Println(Solution([]int{1, 3, 2, 6, 4, 5, 2, 11}) == 10)
	fmt.Println(Solution([]int{1, 3, 2, 6, 4, 5, 2, 11, 2}) == 10)
	fmt.Println(Solution([]int{3, 2, 6, 4, 5, 1, 11}) == 10)
	fmt.Println(Solution([]int{3, 1, 2, 6, 4, 5, 11}) == 10)
	fmt.Println(Solution([]int{0, 1, 1, 0, 1}) == 1)

	fmt.Println(Solution([]int{1, 1, 1, 1, 1}) == 0)
	fmt.Println(Solution([]int{10}) == 0)

	// fmt.Println(Solution([]int{3, 4, 3, 2, 3, -1, 3, 3}) == 15)

	// fmt.Println(Solution([]int{5, 17, 0, 3}) == 17)
	// fmt.Println(Solution([]int{2, 3, 4}) == 0)

	// fmt.Println(Solution([]int{0, 10, -5, -2, 0}) == 10)

	// fmt.Println(Solution([]int{4, 3, 4, 4, 4, 2, 4}) == 0)
	// fmt.Println(Solution([]int{4}) == 0)
	// fmt.Println(Solution([]int{4, 4, 2, 5, 3, 4, 4, 4}) == 0)
	// fmt.Println(Solution([]int{2, 1, 1, 3, 4}) == -1) //  no dominant

	// fmt.Println(Solution([]int{4, 4, -2147483648, 5, 2147483648, 4, 4, 4}) == 0)

}
