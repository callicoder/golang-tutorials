package main
import "fmt"

func main() {
	var BMI = 21.0 
	switch {
		case BMI < 18.5:
			fmt.Println("You're underweight")
		case BMI >= 18.5 && BMI < 25.0:
			fmt.Println("Your weight is normal")
		case BMI >= 25.0 && BMI < 30.0:
			fmt.Println("You're overweight")
		default:
			fmt.Println("You're obese")	
	}
}
