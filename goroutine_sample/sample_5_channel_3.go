package main

import (
	"fmt"
	"time"
)

func producer(start chan bool, out chan int) {
	fmt.Printf("[producer][start]\n")

	<-start // 若in為空，則會一直卡在這，直到有元素投遞進來為止
	fmt.Printf("[producer][in]\n")

	out <- 1
	fmt.Printf("[producer][end]\n")
}

func main() { // main 扮演consumer
	start := make(chan bool, 1)
	out := make(chan int, 1)

	go producer(start, out)

	fmt.Printf("[main][start]\n")
	time.Sleep(time.Second * 3) // 等待3秒，才投遞元素到in裡面，故意讓producer卡一下，才開始做事
	start <- true

	a := <-out
	fmt.Printf("[a][%v]\n", a)
	fmt.Printf("[main][end]\n")
}
