package main

import (
	"fmt"
	"sync"
)

func accumulate(start int, end int, result *int, wg *sync.WaitGroup, mtx *sync.Mutex) {
	defer wg.Done() // 事情做完了，把global counter的值扣1

	for i := start; i < end; i++ {
		mtx.Lock() // 進入critical section

		*result += i

		mtx.Unlock() // 離開critical section
	}
}

func main() {
	result := 0

	wg := &sync.WaitGroup{} // 這是1個global counter，數值是4的意思
	wg.Add(4)

	mtx := &sync.Mutex{}

	go accumulate(1, 25, &result, wg, mtx)   // Σ(1~24)
	go accumulate(25, 50, &result, wg, mtx)  // Σ(25~49)
	go accumulate(50, 75, &result, wg, mtx)  // Σ(50~74)
	go accumulate(75, 101, &result, wg, mtx) // Σ(75~100)

	wg.Wait() // 這行會卡在這邊等，要等到global counter的值歸0時，才可以通過這行

	fmt.Printf("[FINAL RESULT][%d]\n", result) // Σ(1~100) = 5050
}
