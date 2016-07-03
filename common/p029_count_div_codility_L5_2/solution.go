package main

import (
	"fmt"
)

func Solution(A int, B int, K int) int {
	if A == B && (A == K || A == 0) {
		return 1
	}

	if A == B && A != K {
		return 0
	}
	var first int
	if A%K == 0 {
		first = A

	} else {
		first = (A/K + 1) * K
	}

	res := 1 + (B-first)/K
	return res

}

func main() {

	fmt.Println(Solution(6, 11, 2) == 3)
	fmt.Println(Solution(0, 0, 2) == 1)
	fmt.Println(Solution(3, 3, 3) == 1)
	fmt.Println(Solution(6, 12, 2) == 4)
	fmt.Println(Solution(7, 12, 2) == 3)
	fmt.Println(Solution(7, 13, 2) == 3)
	fmt.Println(Solution(10, 10, 5) == 0)
	fmt.Println(Solution(10, 10, 7) == 1)
	fmt.Println(Solution(11, 345, 17) == 20)
	fmt.Println(Solution(2, 16, 3) == 5) // 3,6,9,12,15
	fmt.Println(Solution(5, 16, 3) == 4) // 6,9,12,15
	fmt.Println(Solution(5, 17, 3) == 4) // 6,9,12,15

	// 6 // first
	// (11-6)/2 + 1 = 3

	// 5/3+1 =2
	// 2*3=6 //first value
	// 16-6=10
	// 10/3=3 //=> ans 3+1

	// 6 // first value
	// 17-6=11
	// 11/3=3 // ans => 4

	// 2/3+1=1
	// 1*3=3 // first
	// 16-3=13/3=4
	// 4+1

}
