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

func Solution(A []int) int {
	var s stack

	for _, v := range A {
		if s.Empty() || s.Head() == v {
			s.Push(v)
		} else if s.Head() != v {
			s.Pop()
		}
	}
	if s.Empty() {
		// no leader in slice
		return -1
	}
	cnt, pos := 0, -1
	for i, v := range A {
		if v == s.Head() {
			cnt++
			if pos == -1 {

				pos = i
			}
		}
	}
	if cnt > len(A)/2 {
		return pos
	}
	return -1
}

func main() {
	fmt.Println(Solution([]int{3, 4, 3, 2, 3, -1, 3, 3}) == 0)
	fmt.Println(Solution([]int{0, 1, 1, 0, 1}) == 1)
	fmt.Println(Solution([]int{6, 8, 4, 6, 8, 6, 6}) == 0)
	fmt.Println(Solution([]int{2, 3, 4, 4, 4}) == 2)
	fmt.Println(Solution([]int{4, 3, 4, 4, 4, 2, 4}) == 0)
	fmt.Println(Solution([]int{4}) == 0)
	fmt.Println(Solution([]int{4, 4, 2, 5, 3, 4, 4, 4}) == 0)
	fmt.Println(Solution([]int{2, 1, 1, 3, 4}) == -1) //  no dominant

	fmt.Println(Solution([]int{4, 4, -2147483648, 5, 2147483648, 4, 4, 4}) == 0)

}
