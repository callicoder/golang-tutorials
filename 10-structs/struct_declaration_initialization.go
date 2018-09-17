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
	// Creating an instance of the 'person' struct type
	p1 := Person{"Rajeev", "Singh", 26}
	fmt.Println("Person1: ", p1)

	// Naming fields while initializing the struct
	p2 := Person{
		firstName: "John",
		lastName:  "Snow",
		age:       45,
	}
	fmt.Println("Person2: ", p2)

	// Uninitialized fields are assigned their zero-value
	p3 := Person{
		firstName: "Robert",
	}
	fmt.Println("Person3: ", p3)
}
