package main

import (
	"fmt"
)

func main() {
	fmt.Println("基礎語法:")

	var list []int
	list = make([]int, 4)

	list[0] = 0
	list[1] = 1
	list[2] = 2

	fmt.Printf("%v\n", list) //沒有賦值的部份，會自動填入零值

	fmt.Println("\n測試 slice:")
	list2 := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("%v\n", list2)
	fmt.Printf("%v\n", list2[2:5])

	fmt.Println("\n測試 append:")
	list3 := []int{1, 2}
	list3 = append(list3, 3, 4)
	list3 = append(list3, list2...) // 「...」3個點叫做「解包運算子 (unpack operator)」
	fmt.Printf("%v\n", list3)

	fmt.Println("\n測試空slice:")
	var list4 []int
	fmt.Printf("%#v, %v, %v\n", list4, list4, len(list4))

	list4 = append(list4, 99)
	fmt.Printf("%#v, %v, %v\n", list4, list4, len(list4))
}
