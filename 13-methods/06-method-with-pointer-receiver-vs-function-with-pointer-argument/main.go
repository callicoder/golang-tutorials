package main

import (
	"fmt"
)

type Point struct {
	X, Y float64
}

// Method with Pointer receiver
func (p *Point) Translate(dx, dy float64) {
	p.X = p.X + dx
	p.Y = p.Y + dy
}

// Function with Pointer argument
func TranslateFunc(p *Point, dx, dy float64) {
	p.X = p.X + dx
	p.Y = p.Y + dy
}

func main() {
	p := Point{3, 4}
	ptr := &p
	fmt.Println("Point p = ", p)

	// Calling a Method with Pointer receiver
	p.Translate(2, 6)    // Valid
	ptr.Translate(5, 10) // Valid

	// Calling a Function with a Pointer argument
	TranslateFunc(ptr, 20, 30) // Valid
	// TranslateFunc(p, 20, 30)   // Not Valid
}
