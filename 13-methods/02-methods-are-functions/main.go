package main

import (
	"fmt"
	"math"
)

// Struct type - `Circle`
type Circle struct {
	R float64
}

// Function to calculate the area of Circle
func Area(c Circle) float64 {
	return math.Pi * c.R * c.R
}

/*
Compare the above function with the corresponding method -

func (c Circle) Area() float64 {
	return math.Pi * c.R * c.R; 
}
*/

func main() {
	c := Circle{5.0}

	fmt.Println("Circle : ", c)
	fmt.Println("Area of Circle : ", Area(c))
}
