package main

import (
	"fmt"
	"math"
)

func main() {

	var str []byte
	var row, col int

	fmt.Scanf("%v", &str)

	row = int(math.Floor(math.Sqrt(float64(len(str)))))
	col = int(math.Ceil(math.Sqrt(float64(len(str)))))
	// fmt.Println(len(str), " ", row, " ", col)
	if row*col < len(str) {
		row++
	}
	tmp := make([][]byte, row)
	j := -1
	for i := 0; i < len(str); i++ {
		if i%col == 0 {
			j++
			tmp[j] = make([]byte, col)
		}
		tmp[j][i%col] = str[i]

	}

	res := new([]byte)

	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			if tmp[j][i] != byte(0) {
				*res = append(*res, tmp[j][i])
			}
		}
		*res = append(*res, byte(32))
	}

	fmt.Println(string(*res))
}
