package main

import "fmt"

func main() {
	var a = [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	var slice1 = a[1:]
	var slice2 = a[3:]

	fmt.Println("------- Before Modifications -------")
	fmt.Println("a  = ", a)
	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)

	slice1[0] = "TUE"
	slice1[1] = "WED"
	slice1[2] = "THU"

	slice2[1] = "FRIDAY"

	fmt.Println("\n-------- After Modifications --------")
	fmt.Println("a  = ", a)
	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)
}
