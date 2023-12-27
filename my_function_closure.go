package main

import (
	"fmt"
)

func getFunc() func(int, int) int {
	c := 5

	f := func(a int, b int) int {
		return a + b + c
	}

	return f
}

func main() {
	f := getFunc()

	result := f(2, 3)
	fmt.Printf("[%d][%T]", result, f)
}
