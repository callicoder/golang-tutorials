package main

import "fmt"

func main() {
	// Creates an array of size 10, slices it till index 5, and returns the slice reference
	s := make([]int, 5, 10)
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	/*
		The capacity parameter in the make() function is optional. When omitted, it defaults to the specified length
	*/
	// Creates an array of size 5, and returns a slice reference to it
	var s1 = make([]int, 5)
	fmt.Printf("s1 = %v, len = %d, cap = %d\n", s1, len(s1), cap(s1))

}
