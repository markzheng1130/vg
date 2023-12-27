package main

import (
	"fmt"
	"math"
)

func findMax(nums ...float64) float64 {
	max := math.Inf(-1) // 因為「-inf」格式是float64，格式上為了跟它配套，所以輸入值也要是float64，會比較方便
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}

func main() {
	// 方法1
	max := findMax(5, 3, 1, 8, 99, 4, 2, 6)
	fmt.Printf("%d\n", int(max))

	// 方法2
	list := []float64{5, 3, 1, 8, 99, 4, 2, 6}
	max = findMax(list...) // 透過「...」，達到類似「python」中的「＊args」的效果
	fmt.Printf("%d\n", int(max))
}
