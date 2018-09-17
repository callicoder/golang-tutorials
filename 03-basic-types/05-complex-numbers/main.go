package main

import "fmt"

func main() {
	// === Creating complex numbers ====
	/*
		complex64: both real and imaginary parts are of float32 type.
		complex128: both real and imaginary parts are of float64 type.
	*/
	var x complex64 = 3.4 + 2.9i
	var y = 5 + 7i // Type inferred as `complex128` (default type for complex numbers)

	fmt.Println(x, y)

	// Creating complex no from variables
	var a1 = 4.5
	var a2 = 7.1

	var c = complex(a1, a2) // a1 + a2*i won't work
	fmt.Println(c)

	// ===== Complex No Operations =====
	var a = 3 + 5i
	var b = 2 + 4i

	var res1 = a + b
	var res2 = a - b
	var res3 = a * b
	var res4 = a / b

	fmt.Println(res1, res2, res3, res4)
}
