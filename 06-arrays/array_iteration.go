package main

import "fmt"

func main() {
	// Iterating over and Array and printing its elements
	names := [3]string{"Mark Zuckerberg", "Bill Gates", "Larrt Page"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// Finding the Sum of an Array
	a := [4]float64{3.5, 7.2, 4.8, 9.5}
	sum := float64(0)

	for i := 0; i < len(a); i++ {
		sum = sum + a[i]
	}

	fmt.Printf("Sum of all the elements in array  %v = %f\n", a, sum)
}
