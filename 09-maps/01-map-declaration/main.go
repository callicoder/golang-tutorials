package main

import "fmt"

func main() {
	var m map[string]int
	fmt.Println(m)
	if m == nil {
		fmt.Println("m is nil")
	}

	// The following statement will result in a runtime error
	// m["key"] = 100
}
