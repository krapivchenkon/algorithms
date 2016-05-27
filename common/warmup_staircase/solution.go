package main

import (
	"fmt"
	"strings"
)

func main() {

	// SAMPLE: read from stdin using fmt package
	var N int
	fmt.Scanf("%v", &N)
	// runeChar := []rune("#")[0]
	var tmpl_arr = make([]string, N)
	for i := 0; i < N; i++ {
		tmpl_arr[i] = "#"
	}
	tmpl := strings.Join(tmpl_arr, "")

	var line string
	for i := 0; i < N; i++ {
		line = strings.Replace(tmpl, "#", " ", N-i-1)
		fmt.Println(line)
	}

}
