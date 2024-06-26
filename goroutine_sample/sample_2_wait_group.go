package main

import (
	"fmt"
	"sync"
)

func hello(i int, wg *sync.WaitGroup) {
	fmt.Printf("[hello][%v][say hello]\n", i)
	wg.Done() // 呼叫之後 counter 會減1；且這行等價於「wg.Add(-1)」
}

func main() {
	wg := &sync.WaitGroup{} // 是1個counter
	wg.Add(3)               // 設置counter數值 (要等待的goroutines的數量)

	go hello(1, wg)
	go hello(2, wg)
	go hello(3, wg)

	fmt.Printf("[main][start]\n")
	wg.Wait() // 會卡在這邊，直到「counter 倒數至0」為止

	fmt.Printf("[main][end]\n")
}
