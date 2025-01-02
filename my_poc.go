package main

import (
	"fmt"
)

type myStruct struct {
	nums []string
}

func main() {
	m := map[int]bool{}

	fmt.Printf("[%v]", m[4])
}
