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
		el, err := strconv.ParseInt(s.Text(), 10, 32)
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

func findDigits(v int) int {
	var digits []byte = []byte(strconv.Itoa(v))
	resp := 0
	for _, d := range digits {
		k, _ := strconv.Atoi(string(d))

		if k != 0 && v%k == 0 {
			resp++
		}
	}
	return resp
}

func main() {

	// Setup bufio.Scanner
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)

	//SAMPLE: read integer from stdin
	T := readInt(scan)

	//SAMPLE: read array from the input
	arr := make([]int, T)
	readArray(scan, arr, &T)

	for _, v := range arr {
		fmt.Println(findDigits(v))
	}

}
