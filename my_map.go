package main

import (
	"fmt"
)

func main() {
	d := make(map[string]int)
	d["k1"] = 1
	d["k1"] += 1
	d["k2"] = 2
	fmt.Printf("%v\n", d)

	for key, value := range d {
		fmt.Printf("[%s][%d], ", key, value)
	}

	fmt.Println()

	var ok bool
	_, ok = d["k3"]
	fmt.Printf("%v\n", ok)

	d2 := map[string]int{}
	fmt.Printf("%v, %v\n", d2, d2["k1"])

	counter := map[string]int{} // 1個常見用途，counter
	counter["k1"]++
	counter["k2"]++
	counter["k2"]++
	fmt.Printf("%v\n", counter)

}
