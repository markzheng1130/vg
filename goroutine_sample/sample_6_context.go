package main

import (
	"context"
	"fmt"
	"time"
)

func countNumbers(c context.Context, out chan int) {
	i := 0
	for {
		select {
		case <-c.Done(): // 「c.Done()」會回傳1個「zero size channel」
			out <- i
			return
		default:
			time.Sleep(time.Microsecond * 1000)
			i++
		}
	}
}

func main() {
	out := make(chan int)
	c := context.TODO()                 // 建立1個空context (初始化)
	c1, cancel := context.WithCancel(c) // 轉換成「可取消」型態的context (也可轉換成timeout型態的)

	go countNumbers(c1, out)

	fmt.Printf("[main][start]\n")
	time.Sleep(time.Second * 5) // 等待500毫秒
	cancel()                    // 500毫秒後，呼叫「cancel()」，它會把「c.Done()」中的「zero size channel」給「close掉」

	fmt.Printf("[out][%v]\n", <-out)
	fmt.Printf("[main][end]\n")
}
