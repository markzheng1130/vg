package main

import (
	"fmt"
)

func main() {
	s := "apple"

	switch s {
	case "banana":
		fmt.Printf("Got: %s", s)
	case "orange", "guava", "apple":
		fmt.Printf("Got: %s", s)
	default:
		fmt.Printf("Got: %s", s)
	}
}
