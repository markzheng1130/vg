package main

import (
	"fmt"
)

type Liters float64
type Gallons float64

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func main() {
	carFuel := Gallons(10.0)
	busFuel := Liters(10.0)

	carFuel += ToGallons(busFuel)

	fmt.Printf("[carFuel][%T][%0.2f]]\n", carFuel, carFuel)

	fmt.Printf("[busFuel][%T][%0.2f]]\n", busFuel, busFuel)

	busFuelInGallons := busFuel.ToGallons()
	fmt.Printf("[busFuelInGallons][%T][%0.2f]]\n", busFuelInGallons, busFuelInGallons)
}
