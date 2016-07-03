package main

import (
	"fmt"
)

var proof string = `
https://codesays.com/2014/solution-to-min-avg-two-slice-by-codility/
The key to solve this task is these two patterns:  
(1) There must be some slices, with length of two or three, 
having the minimal average value among all the slices. 
(2) And all the longer slices with minimal average are built 
up with these 2-element and/or 3-element small slices.

Firstly we will prove the statement (1). In all the following 
discussion, we assume the slices have two or more element. In 
every array, there must be at lease one slice, to say S, having 
the Minimal Average (MA). And PLEASE PAY ATTENTION, the following
 proof is done with the S, which HAS the global minimal average.

If the length of S is two or three, it follows our conclusion.
If the length of S is odd, we could divide it into a 3-element 
sub-slice and some 2-element sub-slice. For example, for the 
slice [1, 2, 3, 4, 5], we could get a 3-element sub-slice [1, 2, 3]
 and a 2-element sub-slice [4, 5]. Or differently we could get 
 [1, 2] and [3, 4, 5]. But the split does not matter in the 
 following prove.If the sub-slices have different averages, 
 some of them will have smaller average than MA. But it conflicts
  with the condition that, the MA is known as the global minimal
   average for all slices.  By this contradictory, it’s proved that 
   all the sub-slices MUST have the same average.
If all the sub-slices have the same average, the average of them
 must be MA. So all these sub-slices have overall minimal average
  and length of two or three.
If the length of S is even, we could divide it into some 2-element
 sub-slice. For the slice [1, 2, 3, 4], the only possible split 
 is [1, 2] and [3, 4]. Then, similar with the previous case, all
  these 2-element sub-slices have the global minimal average.
And in the construction in step b, we have already proven the 
statement (2).


`

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func Solution(A []int) int {
	// as minimal avg containes whether in 2 or 3 element
	// slices we only need to iterate over A array once without
	// calculating prefix sums
	pos := 0
	minAvg, curAvg := float64(MaxInt), 0.0

	for i := 0; i < len(A)-2; i++ {
		curAvg = float64((A[i] + A[i+1])) / 2.0
		if curAvg < minAvg {
			minAvg = curAvg
			pos = i
		}
		curAvg = float64((A[i] + A[i+1] + A[i+2])) / 3.0
		if curAvg < minAvg {
			minAvg = curAvg
			pos = i
		}

	}
	// last pair of elements in A checked here
	curAvg = float64((A[len(A)-2] + A[len(A)-1])) / 2.0
	if curAvg < minAvg {
		minAvg = curAvg
		pos = len(A) - 2
	}
	fmt.Println(pos)
	return pos
}

func main() {

	fmt.Println(Solution([]int{4, 2, 2, 5, 1, 5, 8}) == 1)
	fmt.Println(Solution([]int{4, 2}) == 0)
	fmt.Println(Solution([]int{3, 1, 2, 1}) == 1)
	fmt.Println(Solution([]int{-3, 3, -3, 3, -3}) == 0)

}
