package main

import "fmt"

func main() {
	// The zero value of a slice is nil. A nil slice doesn’t have any underlying array, and has a length and capacity of 0
	var s []int
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	if s == nil {
		fmt.Println("s is nil")
	}
}
