package main

import "fmt"

func main() {
	// Creates an array of size 10, slices it till index 5, and returns the slice reference
	var s = make([]int, 5, 10)
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))
}
