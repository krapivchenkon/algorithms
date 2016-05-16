package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func readMatrix(s *bufio.Scanner, arr [][]int32, N *int32) {
	var row, ind int32 = -1, 0
	for s.Scan() {
		if ind%*N == 0 {
			row++
		}
		ind++
		i, err := strconv.ParseInt(s.Text(), 10, 32)
		if err != nil {
			panic(err)
		}
		arr[row] = append(arr[row], int32(i))
	}

}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)
	scan.Scan()
	n, _ := strconv.ParseInt(scan.Text(), 10, 32)
	var N int32 = int32(n)

	var arr = make([][]int32, N, N)

	readMatrix(scan, arr, &N)
	var d1, d2 int32 = 0, 0

	for i := int32(0); i < N; i++ {
		d1 += arr[i][i]
		d2 += arr[i][N-i-1]
	}

	fmt.Println(math.Abs(float64(d1 - d2)))

}
