package main

import (
	"fmt"
	"math"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func Solution(A []int) int {
	sum := 0
	for _, v := range A {
		sum += v
	}
	sum1 := 0
	min := MaxInt
	for i := 0; i < len(A)-1; i++ {
		sum1 += A[i]
		sum -= A[i]
		diff := int(math.Abs(float64(sum1 - sum)))
		if diff == 0 {
			return 0
		}

		if min > diff {
			min = diff
		}

	}
	return min
}

func main() {

	fmt.Println(Solution([]int{3, 1, 2, 4, 3}) == 1)
	fmt.Println(Solution([]int{3, 5}) == 2)

}
