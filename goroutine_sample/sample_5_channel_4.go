package main

import (
	"fmt"
	"time"
)

func consumer(ch chan int) {
	fmt.Printf("[consumer][start]\n")

	for i := range ch { // channel被close時，才有辨法離開此for迴圈
		fmt.Printf("[consumer][%v]\n", i)
	}

	fmt.Printf("[consumer][end]\n")
}

func main() { // main 扮演producer
	ch := make(chan int, 5)

	go consumer(ch)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5

	fmt.Printf("[main][start]\n")
	time.Sleep(time.Second * 5)
	close(ch) // 關閉channel，讓consumer的for迴圈得以離開

	time.Sleep(time.Second * 5) // 確認consumer正常退出之後，再離開main
	fmt.Printf("[main][end]\n")
}
