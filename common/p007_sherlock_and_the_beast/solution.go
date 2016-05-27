package main

import (
	"bufio"
	"errors"
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
	// fmt.Println("changed slice:", arr)
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

func getSplit(v int32) (int32, error) {
	for i := v; i >= 0; i-- {
		if i%3 == 0 && (v-i)%5 == 0 {
			return i, nil
		}
	}
	return -1, errors.New("Split Not found")
}

func findDecent(v int32) (string, error) {
	// res:=
	five := []byte("5")[0]
	three := []byte("3")[0]

	//if divisible by 3 - number can consist only of 5s
	if v%3 == 0 {
		res := make([]byte, v)
		for i, _ := range res {
			res[i] = five
		}
		return string(res), nil
	}
	// find if number of digits can be split in two numbers: one is divisible by 3 adn second is divisible by 5
	split, err := getSplit(v)
	if err != nil {
		return "-1", nil
	}

	res := make([]byte, v)
	for i, _ := range res {
		if i < int(split) {
			res[i] = five
		} else {
			res[i] = three
		}

	}
	return string(res), nil
}

func main() {

	// Setup bufio.Scanner
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)

	//SAMPLE: read integer from stdin
	N := int32(readInt(scan))

	arr := make([]int32, N)
	readArray(scan, arr, &N)

	for _, v := range arr {
		i, err := findDecent(v)
		if err != nil {
			fmt.Println(-1)
		}
		fmt.Println(i)
	}
}
