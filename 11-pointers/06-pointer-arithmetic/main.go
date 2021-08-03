package main

import "fmt"

func main() {
	var x = 67
	var p = &x

	// var p1 = p + 1 // Compiler Error: invalid operation
	var p1 = *p + 1

	fmt.Println(p1)
}
