package main

import (
	"fmt"
)

// helper datastruct in order to solve it with in O(n)
// Another O(n) implementation could be using map and
// counters this is as well O(n) in space complexity as
// a stack solution

type stack []int

func (s stack) Empty() bool { return len(s) == 0 }
func (s stack) Head() int   { return s[len(s)-1] }
func (s *stack) Push(i int) { (*s) = append((*s), i) }
func (s *stack) Pop() int {
	d := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return d
}

func Leader(A []int) int {
	// helper will find leader
	var s stack

	for _, v := range A {
		if s.Empty() || s.Head() == v {
			s.Push(v)
		} else if s.Head() != v {
			s.Pop()
		}
	}
	if s.Empty() {

		return -1
	}
	cnt := 0
	for _, v := range A {
		if v == s.Head() {
			cnt++
		}
	}
	if cnt > len(A)/2 {
		return s.Pop()
	}

	return -1
}

func Solution(A []int) int {
	// first will find leader if it exists in the slice
	// if leaders in two subslices are equal - whole slice will have the same leader
	ld := Leader(A)
	// fmt.Println(A)
	if ld == -1 || len(A) == 1 {
		// no leader in slice
		return 0
	}
	// then we'll find the number it appears in slice
	numLd := 0
	for _, v := range A {
		if v == ld {
			numLd++
		}
	}
	// then we will iterate over A tracking number of
	// occurence of the leader and moving S with every
	// iteration
	left, right, cnt := 0, numLd, 0
	numL, numR := 0, 0

	for i, v := range A {
		if v == ld {
			left++
			right--

		}
		// check if slices satisfy numLeader>n/2
		numL = (i + 1) / 2
		numR = (len(A) - i - 1) / 2
		// fmt.Println(left, numL, right, numR)
		if left > numL && right > numR {
			// leaders on both slices
			// fmt.Println("equi")
			cnt++
		}
	}

	// fmt.Println(cnt)
	return cnt
}

func main() {
	fmt.Println(Solution([]int{1, 5, 2, 1, 4, 0}) == 0)
	fmt.Println(Solution([]int{1, 0, 1, 0, 1}) == 0)
	fmt.Println(Solution([]int{6, 8, 4, 6, 8, 6, 6}) == 0)
	fmt.Println(Solution([]int{4, 3, 4, 4, 4, 2}) == 2)
	fmt.Println(Solution([]int{4, 3, 4, 4, 4, 2, 4}) == 4)
	fmt.Println(Solution([]int{4}) == 0)
	fmt.Println(Solution([]int{4, 4, 2, 5, 3, 4, 4, 4}) == 3)

}
