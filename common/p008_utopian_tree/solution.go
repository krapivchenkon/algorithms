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

func getHeight(v int32) int32 {
	h := int32(1)
	for i := int32(1); i <= v; i++ {
		if i%2 == 0 {
			h++
		} else {
			h += h
		}
	}
	return h
}

func main() {

	// Setup bufio.Scanner
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)

	//SAMPLE: read integer from stdin
	T := int32(readInt(scan))

	//SAMPLE: read array from the input
	arr := make([]int32, T)
	readArray(scan, arr, &T)

	for _, v := range arr {
		fmt.Println(getHeight(v))
	}

}
