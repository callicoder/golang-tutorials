package main

import (
	"fmt"
	"sort"
)

func main() {
	// Sorting a slice of strings by length
	strs := []string{"United States", "India", "France", "United Kingdom", "Spain"}
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	fmt.Println("Sorted strings by length: ", strs)

	// Stable sort
	sort.SliceStable(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	fmt.Println("[Stable] Sorted strings by length: ", strs)

	// Sorting a slice of strings in the reverse order of length
	sort.SliceStable(strs, func(i, j int) bool {
		return len(strs[j]) < len(strs[i])
	})
	fmt.Println("[Stable] Sorted strings by reverse order of length: ", strs)
}
