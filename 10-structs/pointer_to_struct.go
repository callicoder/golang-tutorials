package main

import (
	"fmt"
)

type Student struct {
	rollNumber int
	name       string
}

func main() {
	// instance of student struct type
	s := Student{11, "Jack"}

	// Pointer to a student struct
	ps := &s
	fmt.Println(ps)

	// Accessing struct fields via pointer
	fmt.Println((*ps).name)
	fmt.Println(ps.name) // Same as above: No need to explicitly dereference the pointer

	ps.rollNumber = 31
	fmt.Println(ps)
}
