package main

import "fmt"

func main() {
	var x [5]int // An array of 5 integers

	x[0] = 100
	x[1] = 101
	x[3] = 103
	x[4] = 105

	fmt.Printf("x[0] = %d, x[1] = %d, x[2] = %d\n", x[0], x[1], x[2])
	fmt.Println("x = ", x)
}
