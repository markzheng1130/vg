package main

import (
	"fmt"
)

func main() {
	/// Part-1, 2維陣列
	var list [3][3]string = [3][3]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}}

	fmt.Println("基本印值:")
	fmt.Printf("%v\n", list)
	fmt.Printf("%#v\n", list)

	fmt.Println("\n\n方法1-1:")
	for i := 0; i < len(list); i++ { // 方法1-1
		for j := 0; j < len(list[0]); j++ {
			fmt.Printf("%v, ", list[i][j])
		}
	}

	fmt.Println("\n\n方法2-1:")
	for index, value := range list { // 方法2-1，即是enhanced for loop，用疊代的
		fmt.Printf("[%d]%s, ", index, value)
	}

	fmt.Println("\n\n方法2-2:")
	for row, row_value := range list { // 方法2-2
		for column, _ := range row_value {
			fmt.Printf("%s, ", list[row][column])
		}
	}

	fmt.Println("\n\n方法2-3:")
	for _, row_value := range list { // 方法2-3
		for _, column_value := range row_value {
			fmt.Printf("%s, ", column_value)
		}
	}
}
