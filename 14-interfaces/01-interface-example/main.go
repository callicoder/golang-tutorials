package main

import (
	"fmt"
	"math"
)

// Go Interface - `Shape`
type Shape interface {
	Area() float64
	Perimeter() float64
}

// ==================================================================
// Struct type `Circle` - implements the `Shape` interface implicitly
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Diameter() float64 {
	return 2 * c.Radius
}

// ====================================================================
// Struct type `Rectangle` - implements the `Shape` interface implicitly
type Rectangle struct {
	Length, Width float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

// ==================================================================================
// Generic function to calculate the total area of multiple shapes of different types
func CalculateTotalArea(shapes ...Shape) float64 {
	totalArea := 0.0
	for _, s := range shapes {
		totalArea += s.Area()
	}
	return totalArea
}

// ====================================================================
// Interfaces can also be used as fields
type MyDrawing struct {
	shapes  []Shape
	bgColor string
	fgColor string
}

func (drawing MyDrawing) Area() float64 {
	totalArea := 0.0
	for _, s := range drawing.shapes {
		totalArea += s.Area()
	}
	return totalArea
}

// ==================================================================
func main() {
	var s Shape = Circle{5.0}
	fmt.Printf("Shape Type = %T, Shape Value = %v\n", s, s)
	fmt.Printf("Area = %f, Perimeter = %f\n\n", s.Area(), s.Perimeter())

	var s1 Shape = Rectangle{4.0, 6.0}
	fmt.Printf("Shape Type = %T, Shape Value = %v\n", s1, s1)
	fmt.Printf("Area = %f, Perimeter = %f\n", s1.Area(), s1.Perimeter())

	totalArea := CalculateTotalArea(Circle{2}, Rectangle{4, 5}, Circle{10})
	fmt.Println("Total area = ", totalArea)

	drawing := MyDrawing{
		shapes: []Shape{
			Circle{2},
			Rectangle{3, 5},
			Rectangle{4, 7},
		},
		bgColor: "red",
		fgColor: "white",
	}

	fmt.Println("Drawing", drawing)
	fmt.Println("Drawing Area = ", drawing.Area())

	s = Circle{5}
	fmt.Printf("(%v, %T)\n", s, s)
	fmt.Printf("Shape area = %v\n", s.Area())

	s = Rectangle{4, 7}
	fmt.Printf("(%v, %T)\n", s, s)
	fmt.Printf("Shape area = %v\n", s.Area())

}
