package main

import "fmt"

type Liters float64
type Gallons float64

func main() {
	var carFuel Gallons
	var busFuel Liters
	//convert float64 to Gallons
	carFuel = Gallons(10.0)
	//convert float64 to Liters
	busFuel = Liters(240.0)
	fmt.Println(carFuel, ":", busFuel)
	//auto convert float64 to Gallons
	carFuel = 20.0
	//auto convert float64 to Liters
	busFuel = 100.0
	fmt.Println(carFuel, ":", busFuel)
}
