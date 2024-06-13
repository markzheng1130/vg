package main

import (
	"fmt"
)

func producer(ch chan int) {
	ch <- 10
}

func main() {
	ch := make(chan int, 1)

	go producer(ch)

	i := <-ch // main 扮演consumer
	fmt.Printf("[Got][%v]", i)
}
