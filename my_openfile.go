package main

import (
	"bufio"
	"fmt"
	"os"
)

var FILE_NAME = "sample_data.txt"
var OPTIONS = os.O_RDWR | os.O_APPEND | os.O_CREATE
var FILE_MODE = os.FileMode(0600)

func main() {
	file, err := os.OpenFile(FILE_NAME, OPTIONS, FILE_MODE)

	fmt.Printf("[OPTIONS][%v]\n", OPTIONS)
	fmt.Printf("[FILE_MODE][%v]\n", FILE_MODE)

	if err != nil {
		fmt.Printf("Open file encountered error: %s", err)
	}

	defer file.Close()

	// Part-1, read file.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		fmt.Printf("[%s]\n", line)
	}

	// Part-2, write file.
	s := []byte("zebra\r\n")
	_, err = file.Write(s)

}
