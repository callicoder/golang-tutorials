package main

import "fmt"

func main() {
	var employeeSalary = map[string]float64{
		"John":  78000.00,
		"Steve": 160000.50,
		"David": 85000.00,
	}

	var salary = employeeSalary["Steve"]
	fmt.Printf("Steve's Salary = %f\n", salary)

	// If a key doesn't exist in the map, we get the zero value of the value type
	salary = employeeSalary["Jack"]
	fmt.Printf("Jack's Salary = %f\n", salary)
}
