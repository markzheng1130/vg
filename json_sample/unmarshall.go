package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	LastName  string `json:"lname"` // 這是標籤，用來顯式指定，要匹配JSON中的哪1個欄位
	FirstName string `json:"fname"` // 同上

	Address address // 若結構變數名稱，與JSON變數名稱相同，則可不需標籤

	Age         int    `json:age`       // 若結構中的變數，匹配不到JSON變數，則就不會撈到值，會得到零值
	PhoneNumber string `json:phone_num` // 同上
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
		}
	}`)

	p := person{}
	if err := json.Unmarshal(data, &p); err != nil {
		fmt.Printf("Got error: %v", err)
	}

	fmt.Printf("Got valid JSON, type=%T\n", p)
	fmt.Printf("LastName=%v\n", p.LastName)
	fmt.Printf("FirstName=%v\n", p.FirstName)
	fmt.Printf("Address=%v\n", p.Address)
	fmt.Printf("PhoneNumber=\"%v\"\n", p.PhoneNumber)
	fmt.Printf("Age=%v\n", p.Age)

}
