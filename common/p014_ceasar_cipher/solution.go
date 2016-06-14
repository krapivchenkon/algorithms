package main

import (
	"fmt"
)

const lowerMin = byte(97)
const lowerMax = byte(122)
const upperMin = byte(65)
const upperMax = byte(90)

func main() {

	var N, K int
	var enc []byte
	var char byte
	fmt.Scanf("%v\n%s\n%v", &N, &enc, &K)
	for i, v := range enc {
		if v >= upperMin && v <= upperMax {
			char = v + byte(K)
			for char > upperMax {
				char = char - upperMax + upperMin - 1
			}
			enc[i] = char
		} else if v >= lowerMin && v <= lowerMax {
			char := v + byte(K)
			for char > lowerMax {
				char = char - lowerMax + lowerMin - 1
			}
			enc[i] = char
		}
	}
	fmt.Println(string(enc))

}
