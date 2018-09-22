package main

import "fmt"

func main() {
	var a = 100
	var p = &a

	fmt.Println("a = ", a)
	fmt.Println("p = ", p)
	fmt.Println("*p = ", *p)

	// Changing the value stored in the pointed variable through the pointer
	*p = 2000
	fmt.Println("a (after) = ", a)
}
