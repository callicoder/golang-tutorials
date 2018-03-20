package main

import "fmt"

func main() {
	var src = []string{"Sublime", "VSCode", "IntelliJ", "Eclipse"}
	var dest = make([]string, 2)

	var numElementsCopied = copy(dest, src)

	fmt.Println("src = ", src)
	fmt.Println("dest = ", dest)
	fmt.Println("Number of elements copied from src to dest = ", numElementsCopied)
}
