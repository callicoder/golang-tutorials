package main

import "fmt"

func main() {
	slice1 := []string{"Jack", "John", "Peter"}
	slice2 := []string{"Bill", "Mark", "Steve"}

	slice3 := append(slice1, slice2...)

	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)
	fmt.Println("After appending slice1 & slice2 = ", slice3)
}
