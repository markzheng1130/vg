package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	LastName  string `json:"lname"` // 這是標籤，用來顯式指定，要匹配JSON中的哪1個欄位
	FirstName string `json:"fname"` // 同上

	Address address // 若結構變數名稱，與JSON變數名稱相同，則可不需標籤

	PhoneNumber string `json:phone_num` // 若結構中的變數，匹配不到JSON變數，則就不會撈到值，會得到零值
	Age         int    `json:age`       // 同上
}

type address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	ZipCode int    `json:"zipcode"`
}

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
	}`) // 變數「blood」，因為無法匹配到相同名稱的欄位，所以不會被轉進golang struct

	p := person{}

	if !json.Valid(data) { // 可以單獨驗證，是否為有效JSON格式
		fmt.Printf("Got invalid JSON data")
		os.Exit(1)
	}

	if err := json.Unmarshal(data, &p); err != nil { // 實際進行轉換 JSON to golang struct
		fmt.Printf("Got error: %v", err)
	}

	fmt.Printf("Got valid JSON, type=%T\n", p)
	fmt.Printf("LastName=%v\n", p.LastName)
	fmt.Printf("FirstName=%v\n", p.FirstName)
	fmt.Printf("Address=%+v\n", p.Address)
	fmt.Printf("PhoneNumber=\"%v\"\n", p.PhoneNumber)
	fmt.Printf("Age=%v\n", p.Age)

}
