package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "noName", "Input your name")
	age := flag.Int("age", 0, "Input your age")
	height := flag.Float64("height", 0, "Input your height")
	weight := flag.Float64("weight", 0, "Input your weight")
	flag.Parse()

	fmt.Printf("[%s][%d][%0.2f][%0.2f]\n", *name, *age, *height, *weight)
}

//go run ./my_flag.go -name mark --age 30 -height=170.123 --weight=70.456
//go run ./my_flag.go -h
