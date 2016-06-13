package main

import (
	"fmt"
	"strings"
)

func IndexAll(s, sep string) []int {
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

func findPattern(grid, pattern []string, x, y, xp, yp int) string {

	for i := 0; i <= x-xp; i++ {
		// looking for first line of pattern to match
		// TODO: fix case when pattern present couple times on grid line
	GridIter:
		for _, ind := range IndexAll(grid[i], pattern[0]) {
			// ind := strings.Index(grid[i], pattern[0])
			if ind <= (y - yp) {
				// pattern found on this line
				// check remaining lines
				for j := 1; j < xp; j++ {
					if strings.Compare(grid[i+j][ind:ind+yp], pattern[j]) == 0 {
						continue
					} else {
						continue GridIter
					}
				}
				return "YES"
			}
		}
	}
	return "NO"
}

func main() {

	// SAMPLE: read from stdin using fmt package
	var T, x, y, xp, yp int
	fmt.Scanf("%v\n", &T)
	for t := 0; t < T; t++ {
		fmt.Scanf("%v %v\n", &x, &y)
		grid := make([]string, x)
		for i := 0; i < x; i++ {
			fmt.Scanf("%s\n", &grid[i])
		}
		fmt.Scanf("%v %v", &xp, &yp)
		pattern := make([]string, xp)
		for i := 0; i < xp; i++ {
			fmt.Scanf("%s\n", &pattern[i])
		}

		fmt.Println(findPattern(grid, pattern, x, y, xp, yp))
	}
}
