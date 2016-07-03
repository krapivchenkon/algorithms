package main

import (
	"fmt"
)

func Solution(S string, P []int, Q []int) []int {
	Nucl := map[string]int{
		"A": 1,
		"C": 2,
		"G": 3,
		"T": 4,
	}

	// build consequitive sums for each key
	sums := make(map[string][]int)
	// init pref sums
	sums["A"] = make([]int, len(S)+1)
	sums["C"] = make([]int, len(S)+1)
	sums["G"] = make([]int, len(S)+1)
	sums["T"] = make([]int, len(S)+1)
	// calculate pref sums
	for i, v := range S {
		for k, val := range sums {
			if k == string(v) {
				val[i+1] = val[i] + 1
			} else {
				val[i+1] = val[i]
			}
		}

	}

	fmt.Println(sums)
	// calculate slices for each key
	res := make([]int, len(P))
	for i, v := range P {
		if Q[i] == v {
			// case when slice len is 1
			res[i] = Nucl[string(S[v])]
			continue

		}

		if sums["A"][Q[i]+1]-sums["A"][v] != 0 {
			res[i] = 1
			continue
		}
		if sums["C"][Q[i]+1]-sums["C"][v] != 0 {
			res[i] = 2
			continue
		}
		if sums["G"][Q[i]+1]-sums["G"][v] != 0 {
			res[i] = 3
			continue
		}
		if sums["T"][Q[i]+1]-sums["T"][v] != 0 {
			res[i] = 4
			continue
		}
	}
	fmt.Println(res)
	return res
}

func CmpSlices(A, B []int) bool {
	if len(A) != len(B) {
		return false
	}
	return fmt.Sprintf("%v", A) == fmt.Sprintf("%v", B)
}
func main() {

	fmt.Println(CmpSlices(Solution("CAGCCTA", []int{2, 5, 0}, []int{4, 5, 6}), []int{2, 4, 1}))
	fmt.Println(CmpSlices(Solution("AAAAAAA", []int{2, 5, 0}, []int{4, 5, 6}), []int{1, 1, 1}))
	fmt.Println(CmpSlices(Solution("TTTTTTTT", []int{2}, []int{4}), []int{4}))
	fmt.Println(CmpSlices(Solution("AC", []int{0, 0, 1}, []int{0, 1, 1}), []int{1, 1, 2}))

}
