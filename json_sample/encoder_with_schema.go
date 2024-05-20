package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	LastName  string `json:"lname"` // 這是標籤，用來顯式指定，要匹配JSON中的哪1個欄位
	FirstName string `json:"fname"` // 同上

	Address address // 沒填標籤的話，會採用原變數名稱 (首字大寫)

	phoneNumber string `json:phone_num` // 小寫開頭，無法被轉出
	Age         int    // 沒賦值的話，會以零值被轉出
}

type address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	ZipCode int    `json:"zipcode,omitempty"` // 透過「omitempty」，當變數為零值，或未賦值，就不會被轉出
}

func main() {
	address := address{Street: "da-an Rd", City: "Taipei"}
	p := person{LastName: "Smith", FirstName: "John", Address: address, phoneNumber: "0987-654-321"}

	// --- 差在轉換的語法不一樣 ---
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "\t")
	// --------------------

	if err := encoder.Encode(&p); err != nil {
		fmt.Printf("Got error: %v", err)
	}

}
