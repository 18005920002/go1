package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func main() {
	originalCount := 10
	eatenCount := 4
	fmt.Println("I start by", originalCount, "apples.")
	fmt.Println("Some jerk ate", eatenCount, "apples.")
	fmt.Println("There are", originalCount-eatenCount, "apples.")

	quantity := 4
	length, width := 1.2, 2.4
	customerName := "Damon Cole"
	fmt.Println(customerName)
	fmt.Println("has order", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")

	fmt.Println(reflect.TypeOf(2))
	fmt.Println(reflect.TypeOf(float64(2)))
	height := 3
	fmt.Println("height * width = ", float64(height)*width)
	fmt.Println("height > width ? ", float64(height) > width)
	length = float64(height)
	fmt.Println("height=", height)
	fmt.Println("length=", length)

	var price int = 100
	fmt.Println("Price is", price, "dollars.")
	var taxRate float64 = 0.08
	var tax float64 = float64(price) * taxRate
	fmt.Println("Tax is", tax)
	var total float64 = float64(price) + tax
	fmt.Println("Total cost is", total, "dollars")
	var availableFunds int = 120
	fmt.Println(availableFunds, "dollars available")
	fmt.Println("Within budget?", total < float64(availableFunds))

	var now time.Time = time.Now()
	now.Year()
	fmt.Println(now)
	fmt.Println(reflect.TypeOf(now))
	fmt.Println(now.Year())

	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	fix := replacer.Replace(broken)
	fmt.Println(fix)
}
