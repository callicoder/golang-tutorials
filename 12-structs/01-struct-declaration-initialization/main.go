package main

import (
	"fmt"
)

// Defining a struct type
type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	// Declaring a variable of a `struct` type
	var p Person // // All the struct fields are initialized with their zero value
	fmt.Println(p)

	// Declaring and initializing a struct using a struct literal
	p1 := Person{"Rajeev", "Singh", 26}
	fmt.Println("Person1: ", p1)

	// Naming fields while initializing a struct
	p2 := Person{
		firstName: "John",
		lastName:  "Snow",
		age:       45,
	}
	fmt.Println("Person2: ", p2)

	// Uninitialized fields are set to their corresponding zero-value
	p3 := Person{firstName: "Robert"}
	fmt.Println("Person3: ", p3)
}
