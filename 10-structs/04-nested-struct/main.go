package main

import (
	"fmt"
)

type Address struct {
	houseNo, street, city, state, country string
	zipCode                               int
}

type Customer struct {
	name    string
	age     int
	address Address
}

func main() {
	c := Customer{name: "Rajeev Singh", age: 25}
	c.address = Address{
		houseNo: "747",
		street:  "Golf View Road",
		city:    "Seattle",
		state:   "Washington",
		country: "United States",
		zipCode: 98101,
	}

	fmt.Println("Customer: ", c)
}
