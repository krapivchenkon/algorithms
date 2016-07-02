package main

import (
	"fmt"
)

func Solution(A []int) int {

	perm_sum := len(A) * (len(A) + 1) / 2

	ind := make(map[int]struct{})
	sum := 0
	for _, v := range A {
		if _, ok := ind[v]; !ok && v <= len(A) {
			// key not exists
			sum += v
			ind[v] = struct{}{}
		}
	}
	if sum == perm_sum {
		return 1
	}
	return 0
}

func main() {

	fmt.Println(Solution([]int{4, 1, 3, 2}) == 1)
	fmt.Println(Solution([]int{4, 1, 3}) == 0)

	fmt.Println(Solution([]int{1}) == 1)
	fmt.Println(Solution([]int{3}) == 0)

}
