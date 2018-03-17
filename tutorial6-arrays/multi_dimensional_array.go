package main

import "fmt"

func main() {
	a := [2][2]int{
		{3, 5},
		{7, 9}, // This comma is necessary
	}
	fmt.Println(a)

	// Just like 1D arrays, you don't need to initialize all the elements in a multi-dimensional array.
	// Un-initialized array elements will be assigned the zero value of the array type
	b := [3][4]float64{
		{1, 3},
		{4.5, -3, 7.4, 2},
		{6, 2, 11},
	}
	fmt.Println(b)
}
