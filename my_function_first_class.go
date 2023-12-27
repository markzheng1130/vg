package main

import (
	"fmt"
)

type calculate func(int, int) int

func add(a int, b int) int {
	return a + b
}

func doMathTwice(myMathFunc func(int, int) int, a int, b int) int {
	num := myMathFunc(a, b)
	num += myMathFunc(a, b)
	return num
}

func getMathFunc() func(int, int) int {
	return add
}

func main() {
	// Part-1, 1個function可以被assign給變數
	var myAddFunc func(int, int) int // signature要對齊才行，不然編譯就會報錯 (也可用短變數宣告)
	myAddFunc = add
	fmt.Printf("[%d]\n", myAddFunc(2, 3))

	// Part-2, 1個function可以當做參數，傳入到另外1個funcion中
	fmt.Printf("[%d]\n", doMathTwice(add, 2, 3)) // signature要對齊才行，不然編譯就會報錯

	// Part-3，1個function可以被當做return值，從另1個function中return回來
	var myMathFunc calculate // signature改成參考type def
	myMathFunc = getMathFunc()
	fmt.Printf("[%d]\n", myMathFunc(2, 3))
}
