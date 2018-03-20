package main

import "fmt"

func main() {
	var s []int
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	if s == nil {
		fmt.Println("s is nil")
	}
}
