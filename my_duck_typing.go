package main

import (
	"fmt"
)

type duck struct {
	name string
}

func (d duck) gua(content string) {
	fmt.Printf("[%v][%s]\n", d, content)
}

func (d duck) swim() {
	fmt.Printf("[%v][I can swim!]\n", d)
}

type bmw struct {
	name string
}

func (b bmw) gua(content string) {
	fmt.Printf("[%v][%s]\n", b, content)
}

type gugGuaMaker interface {
	gua(content string)
}

func invokeGua(g gugGuaMaker, content string) {
	g.gua(content)

	if d, ok := g.(duck); ok {
		d.swim()
	} else {
		fmt.Printf("[%v][Can not swim, this is not a duck!]\n", g)
	}
}

func main() {
	var g gugGuaMaker
	g = duck{name: "aDuck"}
	g.gua("duck-gua")

	g = bmw{name: "aBMW"}
	g.gua("BMW-gua")
	fmt.Println("")

	invokeGua(duck{name: "anotherDuck"}, "duck-gua")
	fmt.Println("")

	invokeGua(bmw{name: "anotherBMW"}, "BMW-gua")
}
