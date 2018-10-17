package main

import "fmt"

func main() {
	// Type inference
	var name = "Rajeev Singh" // Type declaration is optional here.
	fmt.Printf("Variable 'name' is of type %T\n", name)

	//================================

	// Multiple variable declarations with inferred types
	var firstName, lastName, age, salary = "John", "Maxwell", 28, 50000.0

	fmt.Printf("firstName: %T, lastName: %T, age: %T, salary: %T\n",
		firstName, lastName, age, salary)
}
