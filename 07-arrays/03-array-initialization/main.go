package main

import "fmt"

func main() {
	// Declaring and initializing an array at the same time
	var a = [5]int{2, 4, 6, 8, 10}
	fmt.Println(a)

	// Short hand declaration for declaring and initializing an array
	b := [5]int{2, 4, 6, 8, 10}
	fmt.Println(b)

	// You don't need to initialize all the elements of the array.
	// The un-initialized elements will be assigned the zero value of the corresponding array type
	c := [5]int{2}
	fmt.Println(c)

	// Letting Go compiler infer the length of the array
	d := [...]int{3, 5, 7, 9, 11, 13, 17}
	fmt.Println(d)
}
