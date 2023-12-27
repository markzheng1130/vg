package main

import (
	"fmt"
	"time"
)

func main() {
	// Part-1，基礎日期格式 (date time structure)
	currentTime := time.Now() // 預設是「local time」
	fmt.Printf("[%v][%v][%v][%v][%v][%v]\n", currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Weekday(), currentTime, currentTime.String())

	currentTimeUTC := time.Now().In(time.UTC) // 也可以選擇是要「UTC+0 time」
	fmt.Printf("[%v][%v][%v][%v][%v]\n", currentTimeUTC.Year(), currentTimeUTC.Month(), currentTimeUTC.Day(), currentTimeUTC.Weekday(), currentTimeUTC)

	year, month, day := currentTime.Date() // 剖析出日期
	fmt.Printf("[%v][%v][%v]\n", year, month, day)

	hour, min, second := currentTime.Clock() // 剖析出時間
	fmt.Printf("[%v][%v][%v]\n\n", hour, min, second)

	// Part-2，測量執行時間的方式
	start := time.Now()
	d := time.Duration(1 * time.Second)
	fmt.Printf("[%T][%v]\n", d, d)
	time.Sleep(d)

	end := time.Now()
	elapsed1 := end.Unix() - start.Unix()
	elapsed2 := end.Sub(start)
	elapsed3 := time.Since(start)
	fmt.Printf("[%T][%v], [%T][%v], [%T][%v]\n", elapsed1, elapsed1, elapsed2, elapsed2, elapsed3, elapsed3)
}
