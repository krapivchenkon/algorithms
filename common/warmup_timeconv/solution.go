package main

import (
	"fmt"
	"time"
)

func main() {

	var a string
	fmt.Scanf("%s", &a)

	const timeFormat = "3:04:05PM"
	t, e := time.Parse(timeFormat, a)
	if e != nil {
		panic(e)
	}
	fmt.Println(t.Format("15:04:05"))
}
