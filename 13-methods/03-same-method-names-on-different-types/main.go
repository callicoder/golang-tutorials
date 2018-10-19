package main

import (
	"fmt"
	"math"
)

type ArithmeticProgression struct {
	A, D float64
}

// Method with receiver `ArithmeticProgression`
func (ap ArithmeticProgression) NthTerm(n int) float64 {
	return ap.A + float64(n-1)*ap.D
}

type GeometricProgression struct {
	A, R float64
}

// Method with receiver `GeometricProgression`
func (gp GeometricProgression) NthTerm(n int) float64 {
	return gp.A * math.Pow(gp.R, float64(n-1))
}

func main() {
	ap := ArithmeticProgression{1, 2} // AP: 1 3 5 7 9 ...
	gp := GeometricProgression{1, 2}  // GP: 1 2 4 8 16 ...

	fmt.Println("5th Term of the Arithmetic series = ", ap.NthTerm(5))
	fmt.Println("5th Term of the Geometric series = ", gp.NthTerm(5))
}
