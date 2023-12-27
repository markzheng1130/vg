package main

import "fmt"

type calculator func(int, int) int

func getCalculator(operator string) calculator {
	switch operator {

	case "+":
		return func(a int, b int) int {
			return a + b
		}
	case "-":
		return func(a int, b int) int {
			return a - b
		}
	default:
		return nil
	}
}

func main() {
	a := 3
	b := 2

	add := getCalculator("+")
	sub := getCalculator("-")

	fmt.Printf("[%d]\n", add(a, b))
	fmt.Printf("[%d]\n", sub(a, b))
}
