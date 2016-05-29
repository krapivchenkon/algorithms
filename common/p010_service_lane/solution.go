package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// HELPER: arr argument should be prealocated with make
func readArray(s *bufio.Scanner, arr []int32, N *int32) {
	for i := int32(0); i < *N && s.Scan(); i++ {
		el, err := strconv.ParseInt(s.Text(), 10, 32)
		if err != nil {
			panic(err)
		}
		arr[i] = int32(el)
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

const MaxUint = ^uint32(0)
const MaxInt = int32(MaxUint >> 1)

func getThroughput(arr []int32, st int32, ex int32) int32 {
	min := MaxInt
	for _, v := range arr[st : ex+1] {
		if v < min {
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
	N := int32(readInt(scan))
	T := readInt(scan)

	//SAMPLE: read array from the input
	arr := make([]int32, N)
	readArray(scan, arr, &N)

	for i := 0; i < T; i++ {
		st := int32(readInt(scan))
		ex := int32(readInt(scan))
		fmt.Println(getThroughput(arr, st, ex))
	}

}
