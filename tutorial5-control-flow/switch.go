package main
import "fmt"

func main() {
	var dayOfWeek = 6
	switch dayOfWeek {
		case 1: fmt.Println("Monday")
		case 2: fmt.Println("Tuesday")
		case 3: fmt.Println("Wednesday")
		case 4: fmt.Println("Thursday")
		case 5: fmt.Println("Friday")
		case 6: {
			fmt.Println("Saturday")
			fmt.Println("Weekend. Yaay!")
		}
		case 7: {
			fmt.Println("Sunday")
			fmt.Println("Weekend. Yaay!")
		}
		default: fmt.Println("Invalid day")
	}
}