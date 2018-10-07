package main

import (
	"fmt"
)

type Student struct {
	RollNumber int
	Name       string
}

func main() {
	// instance of student struct type
	s := Student{11, "Jack"}

	// Pointer to the student struct
	ps := &s
	fmt.Println(ps)

	// Accessing struct fields via pointer
	fmt.Println((*ps).Name)
	fmt.Println(ps.Name) // Same as above: No need to explicitly dereference the pointer

	ps.RollNumber = 31
	fmt.Println(ps)
}
