package main

import "fmt"

func main() {
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
