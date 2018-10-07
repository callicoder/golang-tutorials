package main

import (
	"fmt"
	"github.com/callicoder/golang-tutorials/12-structs/05-struct-exported-vs-unexported/model"
)

func main() {
	c := model.Customer{
		Id: 1, 
		Name: "Rajeev Singh",
	}

	// c.married = true	// Error: can not refer to unexported field or method

	// a := model.address{} // Error: can not refer to unexported name
	
	fmt.Println("Programmer = ", c);
}