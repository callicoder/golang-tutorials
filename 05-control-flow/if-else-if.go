package main
import "fmt"

func main() {
	var BMI = 21.0
	if BMI < 18.5 {
		fmt.Println("You are underweight");
	} else if BMI >= 18.5 && BMI < 25.0 {
		fmt.Println("Your weight is normal");
	} else if BMI >= 25.0 && BMI < 30.0 {
		fmt.Println("You're overweight")
	} else {
		fmt.Println("You're obese")
	}
}