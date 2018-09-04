package main

import "fmt"

func main() {
	var a = 1000
	var p = &a

	fmt.Println("a (before) = ", a)

	// Changing the value stored in the pointed variable through the pointer
	*p = 2000

	fmt.Println("a (after) = ", a)
}
