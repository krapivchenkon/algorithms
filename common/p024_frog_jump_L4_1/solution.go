package main

import (
	"fmt"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func Solution(X int, A []int) int {
	path := X * (X + 1) / 2
	fmt.Println(path)
	ind := make(map[int]struct{})
	curPath := 0
	for i, v := range A {
		if _, ok := ind[v]; !ok && v <= X {
			// key not exists
			curPath += v
			ind[v] = struct{}{}
		}
		if curPath == path {
			return i
		}
	}
	return -1

}

func main() {

	fmt.Println(Solution(5, []int{1, 3, 1, 4, 2, 3, 5, 4}) == 6)
	fmt.Println(Solution(3, []int{1, 3, 1, 4, 2, 3, 5, 4}) == 4)

	fmt.Println(Solution(3, []int{1}) == -1)
	fmt.Println(Solution(1, []int{1}) == 0)

	fmt.Println(Solution(5, []int{1, 3, 1, 4, 2, 3, 3, 4}) == -1)

}
