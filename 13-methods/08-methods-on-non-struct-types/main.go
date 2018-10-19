package main

import (
	"fmt"
)

type MyString string

func (myStr MyString) reverse() string {
	s := string(myStr)
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	myStr := MyString("OLLEH")
	fmt.Println(myStr.reverse())
}
