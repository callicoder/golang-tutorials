package main

import "fmt"

func main() {
	numbers := []float64{3.5, 7.4, 9.2, 5.4}

	sum := 0.0
	for _, number := range numbers {
		sum += number
	}

	fmt.Printf("Total Sum = %.2f\n", sum)
}
