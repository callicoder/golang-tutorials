package main

import "fmt"

func main() {
	var a = [5]string{"C", "C++", "Java", "Python", "Go"}

	var slice1 = a[1:4]
	var slice2 = a[:3]
	var slice3 = a[2:]
	var slice4 = a[:]

	fmt.Println("Array a = ", a)
	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)
	fmt.Println("slice3 = ", slice3)
	fmt.Println("slice4 = ", slice4)
}
