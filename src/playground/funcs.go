package main

import (
	"fmt"
)

func main() {
	area, err := calc(-12.0, 5.3)
	/*if(err!=nil){
		//fmt.Println(err)
		log.Fatal(err)
	}*/
	fmt.Println(err)
	fmt.Printf("need liters: %.2f\n", area)

}

func calc(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	return width * height / 10.0, nil
}
