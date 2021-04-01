package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:]
	fmt.Println(arguments)
	var sum float64 = 0
	for _, arg := range arguments {
		number, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	simpleCount := float64(len(arguments))
	fmt.Printf("avg = %0.2f\n", sum/simpleCount)
	fmt.Printf("max=%0.2f\n", maximum(1.6, 2, 5, 8.6, 3.3, 123.54, 7.2, 3.11, 53.4))
	fmt.Printf("inRange=%0.2f\n", inRange(10, 50, 5, 8.6, 3.3, 123.54, 7.2, 33, 42.6, 3.11, 53.4))
	datas := []float64{53.2, 52.6, 36.3, 64.5}
	fmt.Printf("avg=%0.2f\n", avg(datas...))
}

func maximum(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func inRange(min float64, max float64, numbers ...float64) []float64 {
	filtered := []float64{}
	for _, number := range numbers {
		if number >= min && number <= max {
			filtered = append(filtered, number)
		}
	}
	return filtered
}

func avg(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	simpleCount := float64(len(numbers))

	return sum / simpleCount

}
