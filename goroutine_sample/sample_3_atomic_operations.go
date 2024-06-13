package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func add(a *int32, wg *sync.WaitGroup) {
	for i := int32(1); i <= 10; i++ { // Σ(1~10) = 55
		atomic.AddInt32(a, i) // 用這行的方式，保證加完是「a+55」
		// *a += i // 用這行的方式，不保證加完是「a+55」
	}

	wg.Done()
}

func main() {
	a := int32(0)
	wg := &sync.WaitGroup{}
	wg.Add(5)

	go add(&a, wg)
	go add(&a, wg)
	go add(&a, wg)
	go add(&a, wg)
	go add(&a, wg)

	fmt.Printf("[main][start]\n")
	wg.Wait()
	fmt.Printf("[a][%v]\n", a) // should be 55*5 = 275
	fmt.Printf("[main][end]\n")
}
