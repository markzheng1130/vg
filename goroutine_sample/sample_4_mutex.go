package main

import (
	"fmt"
	"sync"
)

func add(a *int32, wg *sync.WaitGroup, mtx *sync.Mutex) {
	for i := int32(1); i <= 10; i++ { // Σ(1~10) = 55
		mtx.Lock()   // 進入「critical section」
		*a += i      // 保證加完是「a+55」；在鎖的保護下，也可執行atomic沒有支援的，其它種類的運算邏輯
		mtx.Unlock() // 離開「critical section」
	}

	wg.Done()
}

func main() {
	a := int32(0)
	mtx := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(5)

	go add(&a, wg, mtx)
	go add(&a, wg, mtx)
	go add(&a, wg, mtx)
	go add(&a, wg, mtx)
	go add(&a, wg, mtx)

	fmt.Printf("[main][start]\n")
	wg.Wait()

	fmt.Printf("[a][%v]\n", a) // should be 55*5 = 275
	fmt.Printf("[main][end]\n")
}
