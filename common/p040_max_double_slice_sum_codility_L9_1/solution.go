package main

import (
	"fmt"
	"math"
)

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Solution(A []int) int {
	// lets use idea of prefix and suffix sums
	// but we'll right cur Sums in arrays according to Kadane's algorithm

	if len(A) == 3 {
		return 0
	}
	pref := make([]int, len(A)-2) // len-2 as we don't consider edge elements
	suf := make([]int, len(A)-2)

	curPref, curSuf := math.MinInt32, math.MinInt32

	for i := 1; i < len(A)-1; i++ {
		curPref = Max(A[i], curPref+A[i])
		pref[i-1] = curPref

		curSuf = Max(A[len(A)-1-i], curSuf+A[len(A)-1-i])
		suf[len(A)-2-i] = curSuf

	}

	max := math.MinInt32
	// now iterate over pref and suf to find max
	for i := 0; i < len(pref); i++ {

		if i == 0 || i == len(pref)-1 {
			max = Max(suf[1], max)

			max = Max(pref[len(pref)-2], max)

		} else {
			// both pref and suf should be non negative
			max = Max(Max(pref[i-1], 0)+Max(suf[i+1], 0), max)
		}
	}
	if max < 0 {
		return 0
	}
	return max
}

func main() {
	fmt.Println(Solution([]int{3, 2, 6, -1, 4, 5, -1, 2}) == 17)

	fmt.Println(Solution([]int{3, 4, 3, 2, 3, -1, 3, 3}) == 15)
	fmt.Println(Solution([]int{0, 1, 1, 0, 1}) == 2)

	fmt.Println(Solution([]int{5, 17, 0, 3}) == 17)
	fmt.Println(Solution([]int{2, 3, 4}) == 0)
	fmt.Println(Solution([]int{5, 5, 5}) == 0)

	fmt.Println(Solution([]int{0, 10, -5, -2, 0}) == 10)
	// fmt.Println(Solution([]int{4, 3, 4, 4, 4, 2, 4}) == 0)
	// fmt.Println(Solution([]int{4}) == 0)
	// fmt.Println(Solution([]int{4, 4, 2, 5, 3, 4, 4, 4}) == 0)
	// fmt.Println(Solution([]int{2, 1, 1, 3, 4}) == -1) //  no dominant

	// fmt.Println(Solution([]int{4, 4, -2147483648, 5, 2147483648, 4, 4, 4}) == 0)

}
