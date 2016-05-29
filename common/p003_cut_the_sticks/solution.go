package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// HELPER: arr argument should be prealocated with make
func readArray(s *bufio.Scanner, arr []int, N *int) {
	for i := 0; i < *N && s.Scan(); i++ {
		el, err := strconv.ParseInt(s.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		arr[i] = int(el)
	}
}

// HELPER: read in value from the input
func readInt(s *bufio.Scanner) int {
	s.Scan()
	n, err := strconv.ParseInt(s.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	return int(n)
}

const maxUint = ^uint(0)
const maxInt = int(maxUint >> 1)

func Min(arr []int) int {
	min := maxInt
	for _, v := range arr {
		if v != 0 && v < min {
			min = v
		}
	}
	return min

}

func main() {

	// Setup bufio.Scanner
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)

	//SAMPLE: read integer from stdin
	N := readInt(scan)

	//SAMPLE: read array from the input
	arr := make([]int, N)
	readArray(scan, arr, &N)
	left := N

	for left > 0 {
		fmt.Println(left)
		min := Min(arr)
		for i, v := range arr {
			tmp := v - min
			if tmp == 0 {
				left--
			}
			if tmp >= 0 {

				arr[i] = tmp
			}
		}
	}

}
