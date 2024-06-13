package main

import (
	"fmt"
)

func producer(ch chan int) {
	ch <- 1
	ch <- 2
}

func main() { // main 扮演consumer
	ch := make(chan int, 1)

	go producer(ch)

	a := <-ch
	fmt.Printf("[a][%v]\n", a)

	b := <-ch
	fmt.Printf("[b][%v]\n", b)
}
