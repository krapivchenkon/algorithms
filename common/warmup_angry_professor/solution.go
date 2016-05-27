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
	// fmt.Println(changed slice:, arr)
}

func readInt(s *bufio.Scanner) int {
	s.Scan()
	n, err := strconv.ParseInt(s.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	return int(n)

}
func main() {

	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)

	// Number of tests
	T := readInt(scan)

	for t := 0; t < T; t++ {
		// Number of Students
		N := int32(readInt(scan))
		// Threshold value
		K := readInt(scan)

		arr := make([]int32, N)
		readArray(scan, arr, &N)
		onTimeCnt := 0
		for _, v := range arr {
			if v <= 0 {
				onTimeCnt++
			}
		}
		if onTimeCnt >= K {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}

}
