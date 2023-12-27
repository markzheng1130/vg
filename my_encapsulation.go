package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	year  int
	month int
	day   int
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New(fmt.Sprintf("Invalid year: %d", year))
	}

	d.year = year

	return nil
}

func (d *Date) SetMonth(month int) error {
	if (month < 1) || (month > 12) {
		return errors.New(fmt.Sprintf("Invalid month: %d", month))
	}

	d.month = month

	return nil
}

func (d *Date) SetDay(day int) error {
	if (day < 1) || (day > 31) {
		return errors.New(fmt.Sprintf("Invalid day: %d", day))
	}

	d.day = day

	return nil
}

func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}

type Event struct {
	Title string
	Date  // 這是「匿名嵌入」，會讓Date底下的「public method」被升級到跟「Event」同級
}

func main() {
	fmt.Println("Test Date.")
	date := Date{}

	err := date.SetYear(2024)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(1)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(15)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[%v],[%d][%d][%d]\n", date, date.Year(), date.Month(), date.Day())

	fmt.Println("Test Event.")

	event := Event{Title: "My Event"}
	event.SetYear(2024)
	event.SetMonth(2)
	event.SetDay(16)

	fmt.Printf("[%v], [%v][%v][%v]\n", event, event.Year(), event.Month(), event.Day())
}
