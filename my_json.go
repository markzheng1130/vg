package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Part-1, JSON decode.
	// 假裝從外部，收到1個JSON string
	json_input := []byte(`{
		"name": "peter",
		"age": 30,
		"height": 170,
		"contact" : {
		  "address": "taiwan, taipei",
		  "phones": {
			"mobile": "0987-654-321",
			"home": "04-1234-5678"
		  }
		}
	  }`) // 通常在code中，會用「原始字串 (``)」來呈現JSON string

	var m map[string]interface{} // 用1個map來接，這是固定寫法，才有辨法剖析出巢狀結構

	if err := json.Unmarshal(json_input, &m); err != nil {
		fmt.Printf("Error: %v\n\n", err)
	} else {
		fmt.Printf("%v\n\n", m)
	}

	// Part-2, JSON encode.
	indent_two_space := "  " // 設置「2空白」，讓輸出有排版效果(human readable)，也可以用「4空白」，或「\t」
	if json_output, err := json.MarshalIndent(m, "", indent_two_space); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("%T\n\n", json_output)
		fmt.Printf("%s\n\n", string(json_output))
	}
}
