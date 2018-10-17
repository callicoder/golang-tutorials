package main

import "fmt"

func main() {
	s := [][]string{
		{"India", "China"},
		{"USA", "Canada"},
		{"Switzerland", "Germany"},
	}

	fmt.Println("Slice s = ", s)
	fmt.Println("length = ", len(s))
	fmt.Println("capacity = ", cap(s))
}
