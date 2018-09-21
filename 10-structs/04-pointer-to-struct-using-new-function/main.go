package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func main() {
	// You can also get a pointer to a struct using the built-in new() function
	// It allocates enough memory to fit a value of the given struct type, and returns a pointer to it
	pEmp := new(Employee)

	pEmp.id = 1000
	pEmp.name = "Sachin"

	fmt.Println(pEmp)
}
