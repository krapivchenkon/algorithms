package main

import (
	"fmt"
)

func Solution(N int, A []int) []int {
	// write your code in Go 1.4
	arrCnt := make([]int, N)
	// initial max value
	max := 0
	// reset states the level when last time all values were set to max
	reset := 0
	for _, v := range A {
		if v >= 1 && v <= N {

			if arrCnt[v-1] < reset {
				arrCnt[v-1] = reset + 1
			} else {
				arrCnt[v-1] += 1

			}
			// check if new counter value become new max
			if arrCnt[v-1] > max {
				max = arrCnt[v-1]
			}
			continue
		}
		if v == N+1 {
			reset = max

		}
	}
	// fmt.Println(reset, max, arrCnt)
	for i, _ := range arrCnt {
		if arrCnt[i] < reset {
			arrCnt[i] = reset
		}
	}
	return arrCnt

}

func CmpSlices(A, B []int) bool {
	if len(A) != len(B) {
		return false
	}
	return fmt.Sprintf("%v", A) == fmt.Sprintf("%v", B)
}
func main() {

	fmt.Println(CmpSlices(Solution(5, []int{3, 4, 4, 6, 1, 4, 4}), []int{3, 2, 2, 4, 2}))
	fmt.Println(CmpSlices(Solution(2, []int{3, 1, 2}), []int{1, 1}))
	fmt.Println(CmpSlices(Solution(1, []int{1}), []int{1}))
	fmt.Println(CmpSlices(Solution(3, []int{3}), []int{0, 0, 1}))
	// fmt.Println(CmpSlices(Solution([]int{3, 8, 9, 7, 6}), []int{9, 7, 6, 3, 8}))

}
