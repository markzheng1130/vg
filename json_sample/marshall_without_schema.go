package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	address := make(map[string]interface{})
	address["city"] = "Taipei"
	address["street"] = "da-an Rd"
	address["zipcode"] = 106

	p := make(map[string]interface{})
	p["lname"] = "Smith"
	p["fname"] = "John"
	p["address"] = address
	p["blood"] = "O"

	jsonData, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		fmt.Printf("Got error: %v", err)
	}

	fmt.Printf("Got valid json, content=%s", jsonData)
}
