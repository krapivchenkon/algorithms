package main

import (
	"fmt"
	"math"
)

func main() {

	var T int
	var N, C, M float64
	fmt.Scanf("%v", &T)
	for i := 0; i < T; i++ {
		fmt.Scanf("%f %f %f\n", &N, &C, &M)
		// total candies
		res := math.Floor(N / C)
		// wrappers
		rem := res
		for rem >= M {
			can_eat := math.Floor(rem / M)
			res += can_eat
			rem = rem + can_eat*(1-M)
		}

		fmt.Println(res)

	}

}
