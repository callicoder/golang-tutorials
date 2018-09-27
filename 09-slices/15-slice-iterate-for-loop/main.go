package main

import "fmt"

func main() {
	countries := []string{"India", "America", "Russia", "England"}

	for i := 0; i < len(countries); i++ {
		fmt.Println(countries[i])
	}
}
