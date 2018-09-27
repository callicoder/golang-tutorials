package main

import "fmt"

type Point struct {
	x float64
	y float64
}

func main() {
	// Structs are value types.
	p1 := Point{10, 20}
	p2 := p1 // A copy of the struct `p1` is assigned to `p2`
	fmt.Println("p1 = ", p1)
	fmt.Println("p2 = ", p2)

	p2.x = 15
	fmt.Println("\nAfter modifying p2:")
	fmt.Println("p1 = ", p1)
	fmt.Println("p2 = ", p2)
}
