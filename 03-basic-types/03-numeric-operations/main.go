package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b = 4, 5
	var res1 = (a + b) * (a + b) / 2 // Arithmetic operations as int

	a++ // Increment a by 1

	b += 10 // Increment b by 10

	var res2 = a ^ b // Bitwise XOR

	var r = 3.5
	var res3 = math.Pi * r * r // Operations on floating-point type

	var c, d = 4.0, 5.0
	var res4 = (c + d) * (c + d) / 2 // Arithmetic operations as float

	fmt.Printf("res1 : %v, res2 : %v, res3 : %v, res4 : %v\n", res1, res2, res3, res4)
}
