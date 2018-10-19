package main

import (
	"fmt"
)

type Point struct {
	X, Y float64
}

/*
  Translates the current Point, at location (X,Y), by dx along the x axis and dy along the y axis
  so that it now represents the point (X+dx,Y+dy).
*/
func TranslateFunc(p *Point, dx, dy float64) {
	p.X = p.X + dx
	p.Y = p.Y + dy
}

func main() {
	p := Point{3, 4}
	fmt.Println("Point p = ", p)

	TranslateFunc(&p, 7, 9)
	fmt.Println("After Translate, p = ", p)
}
