package main

import (
	"fmt"
)

// Struct type - `Point`
type Point struct {
	X, Y float64
}

func IsAboveFunc(p Point, y float64) bool {
	return p.Y > y
}

/*
Compare the above function with the corresponding method -
func (p Point) IsAbove(y float64) bool {
	return p.Y > y
}
*/

func main() {
	p := Point{2.5, -3.0}

	fmt.Println("Point : ", p)

	fmt.Println("Is Point p located above the line y = 1.0 ? : ", IsAboveFunc(p, 1))
}
