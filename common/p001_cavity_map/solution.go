package main

import (
	"fmt"
)

func DetectCavity(arr [][]byte) {
	rep := []byte("X")[0]

	for i := 1; i < len(arr)-1; i++ {
		for j := 1; j < len(arr)-1; j++ {
			if arr[i-1][j] == rep {
				continue
			}
			if arr[i-1][j] < arr[i][j] &&
				arr[i+1][j] < arr[i][j] &&
				arr[i][j-1] < arr[i][j] &&
				arr[i][j+1] < arr[i][j] {
				arr[i][j] = rep
			}
		}
	}

}

func main() {

	var T int
	var row string
	arr := make([][]byte, T)
	fmt.Scanf("%v\n", &T)
	for t := 0; t < T; t++ {
		fmt.Scanf("%s\n", &row)

		arr = append(arr, []byte(row))
	}

	DetectCavity(arr)

	for _, v := range arr {
		fmt.Println(string(v))
	}

}
