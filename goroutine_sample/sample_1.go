package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("[hello]")
}

func main() {

	go hello()

	for countDown := 3; countDown > 0; countDown-- {
		fmt.Printf("[countDown][%v]\n", countDown)
		time.Sleep(time.Second * 1)
	}
}
