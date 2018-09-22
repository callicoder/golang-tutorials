package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("Original Slice")
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	s = s[1:5]
	fmt.Println("\nAfter slicing from index 1 to 5")
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	s = s[:8]
	fmt.Println("\nAfter extending the length")
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	s = s[2:]
	fmt.Println("\nAfter dropping the first two elements")
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))
}
