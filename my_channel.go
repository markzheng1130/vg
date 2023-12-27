package main

import (
	"fmt"
	"sync"
)

func accumulate(start int, end int, result_collection chan int, wg *sync.WaitGroup) {
	defer wg.Done() // 事情做完了，把global counter的值扣1

	result := 0
	for i := start; i < end; i++ {
		result += i
	}

	result_collection <- result
}

func main() {
	result_collection := make(chan int, 4)

	wg := &sync.WaitGroup{} // 這是1個global counter，數值是4的意思
	wg.Add(4)

	go accumulate(1, 25, result_collection, wg)   // Σ(1~24)
	go accumulate(25, 50, result_collection, wg)  // Σ(25~49)
	go accumulate(50, 75, result_collection, wg)  // Σ(50~74)
	go accumulate(75, 101, result_collection, wg) // Σ(75~100)

	wg.Wait() // 這行會卡在這邊等，要等到global counter的值歸0時，才可以通過這行
	close(result_collection)

	final_result := 0
	for result := range result_collection {
		final_result += result
	}

	fmt.Printf("[FINAL RESULT][%d]\n", final_result) // Σ(1~100) = 5050
}
