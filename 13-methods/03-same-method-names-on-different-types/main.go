package main

import (
	"fmt"
	"math"
)

// Struct type - `Circle`
type Circle struct {
	R float64
}

// Struct type - `Rectangle`
type Rectangle struct {
	Length, Width float64
}

// Method with receiver `Circle`
func (c Circle) Area() float64 {
	return math.Pi * c.R * c.R; 
}

// Method with receiver `Rectangle`
func (r Rectangle) Area() float64 {
	return r.Length * r.Width;
}

func main() {
	c := Circle{5.0}
	r := Rectangle{7.0, 4.0}

	fmt.Println("Circle : ", c)
	fmt.Println("Area of Circle : ", c.Area())

	fmt.Println("Rectangle : ", r)
	fmt.Println("Area of Rectangle : ", r.Area())
}