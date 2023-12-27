package main

import (
	"fmt"
)

func myDeferValue(i int) {
	fmt.Printf("myDeferValue [%d]\n", i)
}

func myDeferPointer(i *int) {
	fmt.Printf("myDeferPointer [%d]\n", *i)
}

func main() {
	i := 10

	defer myDeferValue(i)
	defer myDeferPointer(&i) // defer的項目會推到stack裡面，越晚宣告的會越先執行

	i -= 5
}
