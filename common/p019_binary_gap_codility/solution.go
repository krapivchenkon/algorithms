package main

import (
	"fmt"
	"strconv"
)

func Solution(N int) int {
	var cnt, maxCnt int64
	var st bool

	strN := strconv.FormatInt(int64(N), 2)

	cnt, maxCnt = 0, 0

	for _, v := range strN {

		if string(v) == "1" {
			if st && cnt > 0 {
				if maxCnt < cnt {
					maxCnt = cnt

				}
				cnt = 0

			}
			st = true
		} else {
			if st {
				cnt++
			}
		}
	}

	return int(maxCnt)

}

func main() {

	var N int

	fmt.Scanf("%v\n", &N)

	fmt.Println(Solution(N))

}
