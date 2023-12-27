package main

import (
	"fmt"
)

type person struct {
	name   string
	age    int
	height float64
}

func updatePerson(p *person) { // 傳入的是 struct
	(*p).height += 2
	p.height += 2 // 這與上一行等價
}

func createPerson(name string, age int, height float64) *person { // 回傳的是 struct
	var p person
	p.name = name
	p.age = age
	p.height = height

	fmt.Printf("[memory address in function][%p][%v]\n", &p, p)
	return &p
}

func main() {
	var peter person
	peter.name = "peter"
	peter.age = 40
	peter.height = 176.8

	updatePerson(&peter)
	fmt.Printf("%v\n", peter)

	p := createPerson("mark", 30, 175)
	fmt.Printf("[memory address in main][%p][%v]\n", &(*p), p)

	andy := person{name: "andy", age: 50, height: 190}
	fmt.Printf("%v\n", andy)
}
