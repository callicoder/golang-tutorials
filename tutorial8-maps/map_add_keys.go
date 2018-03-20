package main

import "fmt"

func main() {
	// Initializing a map
	var tinderMatch = make(map[string]string)

	// Adding keys to a map
	tinderMatch["Rajeev"] = "Angelina" // Assigns the value "Angelina" to the key "Rajeev"
	tinderMatch["James"] = "Sophia"
	tinderMatch["David"] = "Emma"

	fmt.Println(tinderMatch)

	/*
	  Adding a key that already exists will simply override
	  the existing key with the new value
	*/
	tinderMatch["Rajeev"] = "Jennifer"
	fmt.Println(tinderMatch)
}
