package main

import (
	"errors"
	"fmt"
	"log"
)

type customError struct {
	s string
}

func (c *customError) Error() string { // 定義1個「to error」的「method」，讓自訂型別的「error」，可以直接被「print」出來
	return fmt.Sprintf("encountered %s", c.s) // error message首字小寫，而且不要有句號
}

func main() {
	// Part-1，自訂error的基礎用法
	err := errors.New("custom error-1")
	fmt.Printf("%s\n\n", err)

	// Part-2，自訂1個error struct
	var ErrCustom error // 用interface來操作
	ErrCustom = &customError{s: "custom error-2"}

	fmt.Printf("%T\n", ErrCustom)

	fmt.Printf("%s\n", ErrCustom.Error())

	fmt.Printf("%s\n", ErrCustom)

	log.Fatal(ErrCustom)
}
