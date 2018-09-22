package main

import "fmt"

func main() {
	// You can create a pointer using the built-in new() function
	ptr := new(int) // Pointer to an int
	*ptr = 100

	fmt.Printf("Ptr = %#x, Ptr value = %d\n", ptr, *ptr)
}
