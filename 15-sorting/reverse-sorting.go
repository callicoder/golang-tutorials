package main

import (
	"fmt"
	"sort"
)

func main() {
	// Sorting a slice of Strings
	strs := []string{"quick", "brown", "fox", "jumpes"}
	sort.Sort(sort.Reverse(sort.StringSlice(strs)))
	fmt.Println("Sorted strings in reverse order: ", strs)

	// Sorting a slice of Integers
	ints := []int{56, 19, 78, 67, 14, 25}
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println("Sorted integers in reverse order: ", ints)

	// Sorting a slice of Floats
	floats := []float64{176.8, 19.5, 20.8, 57.4}
	sort.Sort(sort.Reverse(sort.Float64Slice(floats)))
	fmt.Println("Sorted floats in reverse order: ", floats)
}
