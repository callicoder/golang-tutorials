package main

import (
	"fmt"
)

type Car struct {
	name, model, color string
	weightInKg         float64
}

func main() {
	c := Car{
		name:       "Ferrari",
		model:      "GTC4",
		color:      "Red",
		weightInKg: 1920,
	}

	// Accessing struct fields using dot notation
	fmt.Println("Car name: ", c.name)
	fmt.Println("Car color: ", c.color)

	// Assigning a new value to a struct field
	c.color = "Black"
	fmt.Println("Car: ", c)
}
