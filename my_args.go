package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		fmt.Printf("[%v]\n", arg)
	}
}

// go run my_args.go -name=mark -age=30 -height=170
