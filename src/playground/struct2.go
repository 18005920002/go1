package main

import (
	"com/labs/pg/magazine"
	"fmt"
)

func main() {
	/*address := magazine.Address{
		Street:     "123 Oak St",
		City:       "Omaha",
		State:      "NE",
		PostalCode: "68111",
	}*/
	subscriber := magazine.Subscriber{Name: "ABC"}
	subscriber.Address.Street = "123 Oak St"
	subscriber.Address.City = "Omaha"
	subscriber.Address.State = "NE"
	subscriber.Address.PostalCode = "68111"
	fmt.Println(subscriber.Address)
	fmt.Println("===================")

	fmt.Println(subscriber.Address.Street)
	fmt.Println(subscriber.Address.City)
	fmt.Println(subscriber.Address.State)
	fmt.Println(subscriber.Address.PostalCode)
	fmt.Println("===================")

	fmt.Println(subscriber.Street)
	fmt.Println(subscriber.City)
	fmt.Println(subscriber.State)
	fmt.Println(subscriber.PostalCode)
	fmt.Println("===================")

	employee := magazine.Employee{Name: "ZhangSan"}
	employee.HomeAddress.Street = "456 Elem St"
	employee.HomeAddress.City = "Portland"
	employee.HomeAddress.State = "OR"
	employee.HomeAddress.PostalCode = "97222"
	fmt.Println(employee.HomeAddress.Street)
	fmt.Println(employee.HomeAddress.City)
	fmt.Println(employee.HomeAddress.State)
	fmt.Println(employee.HomeAddress.PostalCode)
}
