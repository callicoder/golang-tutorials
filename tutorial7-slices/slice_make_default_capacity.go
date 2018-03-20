package main

import "fmt"

func main() {
	// Creates an array of size 5, and returns a slice reference to it
	var s = make([]int, 5)
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))
}
