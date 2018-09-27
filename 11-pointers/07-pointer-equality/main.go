package main

import "fmt"

func main() {
	var a = 75
	var p1 = &a
	var p2 = &a

	if p1 == p2 {
		fmt.Println("Both pointers p1 and p2 point to the same variable.")
	}

}
