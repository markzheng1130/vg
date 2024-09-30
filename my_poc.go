package main

import (
	"fmt"
	"reflect"
)

type myStruct struct {
	nums []string
}

func main() {
	s1 := []string{}    // 會得到1個已實現資料結構的空串列
	s2 := []string(nil) // ★ 這個特殊圓括號，會得到nil pointer，尚未實現底層資料結構

	ms := myStruct{}
	s3 := ms.nums // ★ s3會跟s2一樣，會得到nil pointer，尚未實現底層資料結構

	fmt.Printf("[%v][%v][%v][%p][%v]\n", s1, len(s1), reflect.TypeOf(s1), s1, s1)
	fmt.Printf("[%v][%v][%v][%p][%v]\n", s2, len(s2), reflect.TypeOf(s2), s2, s2)
	fmt.Printf("[%v][%v][%v][%p][%v]\n", s3, len(s3), reflect.TypeOf(s3), s3, s3)
}
