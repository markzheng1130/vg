package main

import (
	"fmt"
)

type duck struct {
	name string
}

func (d duck) String() string { // 定義1個「to string」的「method」，讓自訂型別可以直接被「print」出來
	return fmt.Sprintf("This is a duck named: %s", d.name)
}

func main() {
	d := duck{name: "myDuck"}
	fmt.Printf("%s\n", d)
	fmt.Println(d)
}
