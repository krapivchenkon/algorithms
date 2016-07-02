package main

import (
	\"bufio\"
	\"fmt\"
	\"os\"
	\"strconv\"
)

const MaxUint = ^uint(0) 
const MinUint = 0 
const MaxInt = int(MaxUint >> 1) 
const MinInt = -MaxInt - 1


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

// HELPER: find min element in array
func Min(arr []int) int {
	min := maxInt
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min

}

func ShiftSliceLeft(A []int, K int) []int {
	if len(A) == 0 {
		return A
	}
	shift := K % len(A)
	if shift == 0 {
		return A
	}
	res := A[shift:]
	for i, v := range A[:shift] {
		res = append(res, v)
	}
	return res
}

func ShiftSliceRight(A []int, K int) []int {
	if len(A) == 0 {
		return A
	}
	shift := K % len(A)
	if shift == 0 {
		return A
	}
	res := A[len(A)-shift:]
	for _, v := range A[:len(A)-shift] {
		res = append(res, v)
	}
	return res
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

	// SAMPLE: read from stdin using fmt package
	var a, b, res uint32
	fmt.Scanf(\"%v\n%v\", &a, &b)

}
