package main

import "fmt"

func main() {
	var a = [5]string{"Alpha", "Beta", "Gamma", "Delta", "Epsilon"}

	// Creating a slice from the array
	var s []string = a[1:4]

	fmt.Println("Array a = ", a)
	fmt.Println("Slice s = ", s)
}
