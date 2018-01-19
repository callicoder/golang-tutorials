package main
import "fmt"

func main() {
	var a = 3 + 5i
	var b = 2 + 4i

	var res1 = a + b
	var res2 = a - b
	var res3 = a * b
	var res4 = a / b

	fmt.Println(res1, res2, res3, res4)
}