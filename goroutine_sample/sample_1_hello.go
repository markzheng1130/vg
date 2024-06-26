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

	fmt.Printf("[main][start]\n")

	for countDown := 3; countDown > 0; countDown-- { // 給「hello()」3秒的時間，讓它把事情做完 (印出hello)
		fmt.Printf("[countDown][%v]\n", countDown)
		time.Sleep(time.Second * 1)
	}

	fmt.Printf("[main][end]\n")
}
