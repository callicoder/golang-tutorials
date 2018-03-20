package main

import "fmt"

func main() {
	var s = [][]string{
		{"James Smith", "United States"},
		{"Maria Gracia", "England"},
		{"Sarah Johnson", "France"},
	}

	fmt.Println("Slice s = ", s)
	fmt.Println("length = ", len(s))
	fmt.Println("capacity = ", cap(s))
}
