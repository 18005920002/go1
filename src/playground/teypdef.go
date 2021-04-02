package main

import (
	"fmt"
)

type Liters float64
type Milliliters float64
type Gallons float64

func main() {
	var carFuel Gallons
	var busFuel Liters
	//convert float64 to Gallons
	carFuel = Gallons(1.2)
	//convert float64 to Liters
	busFuel = Liters(4.5)
	fmt.Println(carFuel, ":", busFuel)
	//auto convert float64 to Gallons
	//carFuel = 10.0
	//auto convert float64 to Liters
	//busFuel = 240.0
	//fmt.Println(carFuel, ":", busFuel)

	carFuel += ToGallons(Liters(40))
	busFuel += toLiters(Gallons(30))
	fmt.Printf("%0.1f:%0.1f", carFuel, busFuel)

}

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}
func toLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}
