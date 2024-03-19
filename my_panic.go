package main

import (
	"fmt"
)

func myFunc(i int) {
	if i == 0 {
		panic("Panic happens, count down to zero!")
	} else {
		fmt.Printf("myFunc[%d]\n", i)
		myFunc(i - 1)
	}
}

func myErrorHandling() {
	fmt.Printf("[ENTER][ERROR HANDLING]\n")

	if r := recover(); r != nil {
		fmt.Printf("[%s]\n", r) //「recover()」要躲在「defer function」裡面，才會有機會被執行到
	}

	fmt.Printf("[END][ERROR HANDLING]\n")
}

func main() {
	defer myErrorHandling()
	myFunc(3)
}
