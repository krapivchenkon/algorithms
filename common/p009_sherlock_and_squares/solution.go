package main

import (
	"fmt"
	"math"
)

func main() {

	var T, a, b int
	fmt.Scanf("%v\n", &T)
	for i := 0; i < T; i++ {
		res := 0
		fmt.Scanf("%v %v\n", &a, &b)
		sq := math.Sqrt(float64(a))
		sq_floor := math.Floor(sq)
		if sq_floor-sq == float64(0) {
			// lower boundary is a square
			res++
		}
		for true {
			sq_floor++
			if math.Pow(sq_floor, 2) <= float64(b) {
				res++
			} else {
				break
			}
		}
		fmt.Println(res)

	}

}
