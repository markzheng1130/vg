package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := []byte(`{
	"lname": "Smith",
	"fname": "John",
	"address": {
		"street": "da-an Rd",
		"city": "Taipei",
		"zipcode": 106
	},
	"blood":"O"
	}`)

	var p map[string]interface{}

	json.Unmarshal(data, &p)
	fmt.Printf("Got valid JSON, p=%v\n", p)
}
