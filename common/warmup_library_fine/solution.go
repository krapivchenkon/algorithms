package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	const dateFmt = "2 1 2006"

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	dtActStr := scanner.Text()
	scanner.Scan()
	dtDueStr := scanner.Text()

	dtAct, e := time.Parse(dateFmt, dtActStr)
	if e != nil {
		panic(e)
	}
	dtDue, e := time.Parse(dateFmt, dtDueStr)
	if e != nil {
		panic(e)
	}

	if dtAct.Before(dtDue) {
		fmt.Println(0)
		return
	}

	if dtAct.Year() == dtDue.Year() {
		if dtAct.Month() == dtDue.Month() {
			fmt.Println(15 * (dtAct.Day() - dtDue.Day()))
		} else {
			// fmt.Println("Test", int(dtAct.Month()), "   ", dtDue.Month())
			fmt.Println(500 * (int(dtAct.Month()) - int(dtDue.Month())))
			// fmt.Println("Test1")
		}
	} else {
		fmt.Println(10000)
	}

}
