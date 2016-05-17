package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// arr argument should be prealocated with make
func readArray(s *bufio.Scanner, arr []int32, N *int32) {
	for i := int32(0); i < *N && s.Scan(); i++ {
		el, err := strconv.ParseInt(s.Text(), 10, 32)
		if err != nil {
			panic(err)
		}
		arr[i] = int32(el)
	}
	// fmt.Println("changed slice:", arr)
}

func main() {

	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)
	scan.Scan()
	n, _ := strconv.ParseInt(scan.Text(), 10, 32)
	var N int32 = int32(n)

	arr := make([]int32, N)
	readArray(scan, arr, &N)

	var sum int32
	for _, v := range arr {
		sum += v
	}
	fmt.Println(sum)
}
