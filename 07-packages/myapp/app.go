package main

import (
	"fmt"
	"github.com/callicoder/golang-tutorials/07-packages/numbers"
	"github.com/callicoder/golang-tutorials/07-packages/strings"	
	"github.com/callicoder/golang-tutorials/07-packages/strings/greeting" // Importing a nested package
	str "strings"	// Package Alias
)

func main() {
	fmt.Println(numbers.IsPrime(19))

	fmt.Println(greeting.WelcomeText)

	fmt.Println(strings.Reverse("callicoder"))

	fmt.Println(str.Count("Go is Awesome. I love Go", "Go"))
}