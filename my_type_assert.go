package main

import (
	"fmt"
)

type duck struct {
	name string
}

func main() {
	// Part-1，基本款 type assert
	var s interface{} = "abc"
	fmt.Printf("[%T][%T][%s]\n\n", s, s.(string), s.(string))

	// Part-2，type switch (特殊語法)
	collection := []interface{}{1, 2.3, "456", duck{name: "d"}}

	for _, item := range collection {
		switch instance := item.(type) {

		case duck:
			fmt.Printf("[%T][%s]\n", instance, instance.name)
		case string:
			fmt.Printf("[%T][%s]\n", instance, instance)
		case float64:
			fmt.Printf("[%T][%0.5f]\n", instance, instance)
		case int:
			fmt.Printf("[%T][%d]\n", instance, instance)
		}
	}
}
