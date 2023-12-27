package main

import (
	"fmt"
)

type person struct {
	name   string
	age    int
	height int
}

func (p *person) String() string {
	return fmt.Sprintf("My name is %s, %d ages old, and height %d.", p.name, p.age, p.height)
}

func main() {
	p := &person{name: "peter", age: 30, height: 180}
	fmt.Printf("[%T][%v]", p, p)
}
