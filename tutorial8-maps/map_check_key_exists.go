package main

import "fmt"

func main() {
	var employees = map[int]string{
		1001: "Rajeev",
		1002: "Sachin",
		1003: "James",
	}

	printEmployee(employees, 1001)
	printEmployee(employees, 1010)

	if isEmployeeExists(employees, 1002) {
		fmt.Println("EmployeeId 1002 found")
	}
}

func printEmployee(employees map[int]string, employeeId int) {
	if name, ok := employees[employeeId]; ok {
		fmt.Printf("name = %s, ok = %v\n", name, ok)
	} else {
		fmt.Printf("EmployeeId %d not found\n", employeeId)
	}
}

func isEmployeeExists(employees map[int]string, employeeId int) bool {
	_, ok := employees[employeeId]
	return ok
}
