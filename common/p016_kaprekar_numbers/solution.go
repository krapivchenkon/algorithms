package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var a, b int
	var res []string
	fmt.Scanf("%v\n%v", &a, &b)

	for i := a; i <= b; i++ {
		n := len(strconv.Itoa(i))
		s := strconv.Itoa(i * i)
		r, err := strconv.Atoi(s[len(s)-n:])
		if err != nil {
			r = 0
		}
		l, err := strconv.Atoi(s[:len(s)-n])
		if err != nil {
			l = 0
		}

		if (r + l) == i {
			res = append(res, strconv.Itoa(i))
		}
	}

	if len(res) == 0 {
		fmt.Println("INVALID RANGE")
	} else {

		fmt.Println(strings.Join(res, " "))
	}

}
