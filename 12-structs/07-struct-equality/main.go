package main

import "fmt"

type Point struct {
	X float64
	Y float64
}

func main() {
	// Two structs are equal if all their corresponding fields are equal.
	p1 := Point{3.4, 5.2}
	p2 := Point{3.4, 5.2}

	if p1 == p2 {
		fmt.Println("Point p1 and p2 are equal.")
	} else {
		fmt.Println("Point p1 and p2 are not equal.")
	}
}
