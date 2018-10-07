package main

import (
	"math"
	"fmt"
)

type Circle struct {
	R float64
}

// Method with value receiver
func (c Circle) Area() float64 {
	return math.Pi * c.R * c.R; 
}

// Function with value argument
func AreaFunc(c Circle) float64 {
	return math.Pi * c.R * c.R
}

func main() {
	c := Circle{5}	
	ptr := &c

	fmt.Println("Circle c = ", c)

	// Calling a Method with Value receiver
	c.Area()		// Valid
	ptr.Area()	    // Valid

	// Calling a Function with a Value argument
	AreaFunc(c)     // Valid
	// AreaFunc(ptr)   // Not Valid
}