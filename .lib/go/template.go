package main

import (
	\"bufio\"
	\"fmt\"
	\"os\"
	\"strconv\"
)

// HELPER: function reads 2-dimensional matrix into slice of slices
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

// HELPER: arr argument should be prealocated with make
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

	// SAMPLE: read input sample using bufio Scanner
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)
	scan.Scan()
	n, _ := strconv.ParseInt(scan.Text(), 10, 32)
	var N int32 = int32(n)

	// SAMPLE: read from stdin using fmt package
	var a, b, res uint32
	fmt.Scanf("%v\n%v", &a, &b)

}
