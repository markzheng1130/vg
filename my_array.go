package main

import (
	"fmt"
)

func main() {
	// Part-1，雜項
	fmt.Println("\n\n方法-1:")
	var list2 [3]int = [3]int{1, 2, 3} // 測試初始方式
	fmt.Printf("%v\n", list2)

	var a = [3]int{1, 2, 3}
	var b = [3]int{1, 2, 3}
	var c = [3]int{1, 2, 4}
	fmt.Printf("[%v][%p][%p]\n", a == b, &a, &b) // array可以直接比較，是比裡面的元素
	fmt.Printf("[%v]\n", b == c)
}
