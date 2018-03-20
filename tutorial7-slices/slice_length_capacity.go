package main

import "fmt"

func main() {
	var a = [6]int{10, 20, 30, 40, 50, 60}
	var s = a[1:4]

	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))
}
