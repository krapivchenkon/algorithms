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

func main() {

	// Setup bufio.Scanner
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)

	//SAMPLE: read integer from stdin
	n := readInt(scan)
	page := readInt(scan)
	//SAMPLE: read array from the input
	arr := make([]int, n)
	readArray(scan, arr, &n)

	var curPage, magic int = 1, 0
	for i := 0; i < n; i++ {

		chapt := arr[i] / page
		if arr[i]%page != 0 {
			chapt++
		}

		start, end := 1, page
		for j := 0; j < chapt; j++ {
			start = 1 + j*page
			end = page + j*page
			if end > arr[i] {
				end = arr[i]
			}
			if curPage >= start && curPage <= end {
				magic++
			}
			curPage++

		}
	}
	fmt.Println(magic)

}
