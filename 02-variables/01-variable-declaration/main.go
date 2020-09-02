package main

import "fmt"

func main() {
	// Declaring Variables
	var myStr string = "Hello"
	var myInt int = 100
	var myFloat float64 = 45.12
	fmt.Println(myStr, myInt, myFloat)

	//================================

	// Multiple Declarations
	var (
		employeeId          int    = 5
		firstName, lastName string = "Justin", "Jones"
	)
	fmt.Println(employeeId, firstName, lastName)

	//================================

	// Short variable declaration syntax
	name := "Rajeev Singh"
	age, salary, isProgrammer := 55, 50000.0, true

	fmt.Println(name, age, salary, isProgrammer)
}
