package main

import "fmt"

func main() {
	/*
		Maps are reference types. When you assign a map to another variable, they both refer to the same underlying data structure.
		Therefore changes done by one variable will be visible to the other
	*/
	var m1 = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}

	var m2 = m1
	fmt.Println("m1 = ", m1)
	fmt.Println("m2 = ", m2)

	m2["ten"] = 10
	fmt.Println("\nm1 = ", m1)
	fmt.Println("m2 = ", m2)
}
