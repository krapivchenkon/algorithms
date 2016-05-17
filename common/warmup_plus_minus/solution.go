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

func main() {

	// SAMPLE: read input sample using bufio Scanner
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)
	scan.Scan()
	n, _ := strconv.ParseInt(scan.Text(), 10, 32)
	N := int32(n)
	arr := make([]int32, N)

	readArray(scan, arr, &N)
	var neg, pos, zer float64 = 0, 0, 0
	for _, v := range arr {
		switch {
		case v == 0:
			zer++
		case v < 0:
			neg++
		default:
			pos++
		}
	}

	fmt.Printf("%.6f\n", pos/float64(N))
	fmt.Printf("%.6f\n", neg/float64(N))
	fmt.Printf("%.6f\n", zer/float64(N))

}
