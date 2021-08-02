package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UsersByAge []User

func (u UsersByAge) Len() int {
	return len(u)
}
func (u UsersByAge) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
func (u UsersByAge) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

func main() {
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

	// Sorting a slice of users by age
	sort.Sort(UsersByAge(users))
	fmt.Println("Sorted users by age: ", users)

	// Sorting a slice of strings in the reverse order of length
	sort.Stable(UsersByAge(users))
	fmt.Println("[Stable] Sorted users by age: ", users)

}
