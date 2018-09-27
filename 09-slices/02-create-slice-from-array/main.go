package main

import "fmt"

func main() {
	var a = [5]string{"Alpha", "Beta", "Gamma", "Delta", "Epsilon"}

	// Creating a slice from the array
	var s []string = a[1:4]

	fmt.Println("Array a = ", a)
	fmt.Println("Slice s = ", s)

	/*
		low and high parameters are optional in a[low:high]
		The default value for low is 0, and high is the length of the slice.
	*/
	slice1 := a[1:4]
	slice2 := a[:3]
	slice3 := a[2:]
	slice4 := a[:]

	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)
	fmt.Println("slice3 = ", slice3)
	fmt.Println("slice4 = ", slice4)
}
