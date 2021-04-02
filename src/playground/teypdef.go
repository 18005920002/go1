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
	fmt.Printf("%0.1f:%0.1f\n", carFuel, busFuel)

	soda := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())
	water := Milliliters(500)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", water, water.ToGallons())
	milk := Gallons(2)
	fmt.Printf("%0.3f gallons equals %0.3f liters\n", milk, milk.ToLiters())
	fmt.Printf("%0.3f gallons equals %0.3f milliliters\n", milk, milk.ToMilliliters())

}

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}
func toLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func (m Liters) ToGallons() Gallons {
	return Gallons(m * 0.264)
}
func (m Milliliters) ToGallons() Gallons {
	return Gallons(m / 1000 * 0.264)
}
func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}
func (g Gallons) ToMilliliters() Milliliters {
	return Milliliters(g * 3785.41)
}
