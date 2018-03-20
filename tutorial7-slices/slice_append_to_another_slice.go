package main

import "fmt"

func main() {
	var slice1 = []string{"Jack", "John", "Peter"}
	var slice2 = []string{"Bill", "Mark", "Steve"}

	var slice3 = append(slice1, slice2...)

	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)
	fmt.Println("After appending slice1 & slice2 = ", slice3)
}
