package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var shutdownFlag bool = false

func main() {
	fmt.Printf("[START][MAIN]\n")
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go signalHandler(sigs)

	counter := 0
	for !shutdownFlag {
		fmt.Printf("[%d][%v]\n", counter, shutdownFlag)

		if counter == 5 {
			fmt.Printf("[TRIGGERING][SIGTERM]\n")
			sigs <- syscall.SIGTERM
		}

		time.Sleep(1 * time.Second)
		counter++
	}

	fmt.Printf("[%d][%v]\n", counter, shutdownFlag)
	fmt.Printf("[END][MAIN]\n") // 這行預期要執行不到
}

func signalHandler(sigs chan os.Signal) {
loopSignalHandler:
	for true {
		sig := <-sigs
		switch sig {

		case syscall.SIGTERM:
			fmt.Printf("[RECEIVED][SIGTERM], sleep a while..\n")
			time.Sleep(5 * time.Second) // 雖然收到shutdown signal了，但故意讓main再繼續跑一下
			shutdownFlag = true
			break loopSignalHandler

		case syscall.SIGINT:
			fmt.Printf("[RECEIVED][SIGINT], sleep a while..\n")
			time.Sleep(5 * time.Second)
			shutdownFlag = true
			break loopSignalHandler

		}
	}
}
