package main

import (
	"testing"
	"vg/hello"
)

func TestHello(t *testing.T) {
	// Arrange
	name := "mark"

	// Act
	actual := hello.Hello(name)

	// Assert
	expect := "hello, world, Eric"

	if actual != expect {
		t.Errorf("[actual][%s][expect][%s]", actual, expect)
	}
}

// 執行: go test my_hello_test.go -v
