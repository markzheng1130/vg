package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1) // 只有1個goroutine，也可以使用channel，這是最簡單的範例

	ch <- 1   // 透過「<-」傳進進channel
	i := <-ch // 透過「<-」從channel倒值出來

	fmt.Printf("[Got][%v]", i)
}
