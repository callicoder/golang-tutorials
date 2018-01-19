package main
import "fmt"

func main() {
	var myInt8 int8 = 97

	/*
	  When you don't declare any type explicitly, the type inferred is `int`
	  (The default type for integers)
	*/
	var myInt = 1200

	var myUint uint = 500

	var myHexNumber = 0xFF  // Use prefix '0x' or '0X' for declaring hexadecimal numbers
	var myOctalNumber = 034 // Use prefix '0' for declaring octal numbers

	fmt.Printf("%d, %d, %d, %#x, %#o\n", myInt8, myInt, myUint, myHexNumber, myOctalNumber)
}