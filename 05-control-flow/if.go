package main
import "fmt"

func main() {
	// If Statement
	var x = 25
	if(x % 5 == 0) {
		fmt.Printf("%d is a multiple of 5\n", x)
	}

	// Parentheses are Optional
	var y = -1
	if y < 0 {
		fmt.Printf("%d is negative\n", y)
	}
	
	// If with a condition consisting of short circuit operators
	var age = 21
	if age >= 17 && age <= 30 {
		fmt.Println("My Age is between 17 and 30")
	}	

	// If with a short statement	
	if n := 10; n%2 == 0 {
		fmt.Printf("%d is even\n", n)
	} 
	
}