package main

import (
	// "bytes"
	"fmt"
	"strings"
)

func IndexAll(s, sep string) []int {
	// arr := []byte(s)
	// sepb := []byte(sep)
	var indeces []int
	i := 0
	for {
		ind := strings.Index(s[i:], sep)
		if ind >= 0 {
			indeces = append(indeces, ind+i)
			i = i + ind + 1
		} else {
			break
		}

	}
	return indeces
}

func main() {

	fmt.Println(strings.Count("12121212", "1212"))

	fmt.Println(IndexAll("12121212", "1212"))

	fmt.Println(IndexAll("12121212", "542"))
}
