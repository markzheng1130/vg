package main

import (
	"fmt"
)

type num = int

func main() {
	var myNum = num(10)
	anotherNum := 20

	fmt.Printf("[%d][%T][%T]", myNum+anotherNum, myNum, anotherNum)
}
