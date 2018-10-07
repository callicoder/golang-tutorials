package main

import (
	"fmt"
	"math"
)

// Struct type - `Circle`
type Circle struct {
	R float64
}

// Method with receiver `Circle`
func (c Circle) Area() float64 {
	return math.Pi * c.R * c.R; 
}

func main() {
	c := Circle{5.0}

	fmt.Println("Circle : ", c)
	fmt.Println("Area of Circle : ", c.Area())
}