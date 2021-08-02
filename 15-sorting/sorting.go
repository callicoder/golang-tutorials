package main

import (
	"fmt"
	"sort"
)

func main() {
	// Sorting a slice of Strings
	strs := []string{"quick", "brown", "fox", "jumpes"}
	sort.Strings(strs)
	fmt.Println("Sorted strings: ", strs)

	// Sorting a slice of Integers
	ints := []int{56, 19, 78, 67, 14, 25}
	sort.Ints(ints)
	fmt.Println("Sorted integers: ", ints)

	// Sorting a slice of Floats
	floats := []float64{176.8, 19.5, 20.8, 57.4}
	sort.Float64s(floats)
	fmt.Println("Sorted floats: ", floats)
}
