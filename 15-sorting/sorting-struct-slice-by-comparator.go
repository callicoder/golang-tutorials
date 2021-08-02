package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

func main() {
	// Sorting a slice of structs by a field
	users := []User{
		{
			Name: "Rajeev",
			Age:  28,
		},
		{
			Name: "Monica",
			Age:  31,
		},
		{
			Name: "John",
			Age:  56,
		},
		{
			Name: "Amanda",
			Age:  16,
		},
		{
			Name: "Steven",
			Age:  28,
		},
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].Age < users[j].Age
	})
	fmt.Println("Sorted users by age: ", users)

	// Stable sort
	sort.SliceStable(users, func(i, j int) bool {
		return users[i].Age < users[j].Age
	})
	fmt.Println("Sorted users by age: ", users)
}
