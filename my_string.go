package main

import (
	"fmt"
)

func main() {
	s := "abcde"
	for _, value := range s {
		fmt.Printf("[%s][%T]\n", string(value), value)
	}
}
