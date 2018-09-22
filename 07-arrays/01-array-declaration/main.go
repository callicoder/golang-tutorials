package main

import "fmt"

func main() {
	var x [5]int // An array of 5 integers
	fmt.Println(x)

	var y [8]string // An array of 8 strings
	fmt.Println(y)

	var z [3]complex128 // An array of 3 complex numbers
	fmt.Println(z)

	// By default, all the array elements are assigned the zero value of the array type.
	// For example, if we declare an integer array, all the elements will be initialized with zero.
	// If we declare a string array, all the elements will be initialized with an empty string, and so on.
}
