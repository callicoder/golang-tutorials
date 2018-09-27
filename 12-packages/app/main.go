package main

import (
	"fmt"
	"github.com/callicoder/golang-tutorials/12-packages/numbers"
	"github.com/callicoder/golang-tutorials/12-packages/strings"
	str "strings"
)

func main() {
	fmt.Println(strings.WelcomeText)
	fmt.Println(strings.Reverse("rajeev"))
	fmt.Println(numbers.IsPrime(9))
	fmt.Println(str.Contains("Hello, World", "Wo"))
}