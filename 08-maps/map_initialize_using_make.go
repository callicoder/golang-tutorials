package main

import "fmt"

func main() {
	var m = make(map[string]int)

	fmt.Println(m)

	if m == nil {
		fmt.Println("m is nil")
	} else {
		fmt.Println("m is not nil")
	}

	// make() function returns an initialized and ready to use map.
	// Since it is initialized, you can add new keys to it.
	m["one hundred"] = 100
	fmt.Println(m)

}
