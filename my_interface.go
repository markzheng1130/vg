package main

import "fmt"

type cat struct {
}

type dog string

type person struct {
	name string
}

func (c cat) Speak() string {
	return fmt.Sprintf("Cat speak.")
}

func (d dog) Speak() string {
	return fmt.Sprintf("Dog speak.")
}

func (p person) Speak() string {
	return fmt.Sprintf("Person speak. (%s)", p.name)
}

type Speaker interface {
	Speak() string
}

func DoSpeak(s Speaker) {
	fmt.Printf("[%T][%s]\n", s, s.Speak())
}

func main() {
	c := cat{}
	d := dog("")
	p := person{name: "mark"}

	var s Speaker
	s = person{name: "peter"}

	DoSpeak(c)
	DoSpeak(d)
	DoSpeak(p)
	DoSpeak(s)
}
