package main

import (
	"fmt"
)

func main() {
	a := 10

	// 有使用到「initial state」的if-else
	if count := a; count > 5 {
		fmt.Printf("True, got: %d", a)
	} else {
		fmt.Printf("False, got: %d", a)
	}
}
