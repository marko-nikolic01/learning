package main

import (
	"fmt"
	"math"
)

type Triangle struct {
	x1 float64
	y1 float64

	x2 float64
	y2 float64

	x3 float64
	y3 float64
}

func main() {
	triangle := Triangle{0, 0, 1, 0, 0, 1}

	p := calculateArea(triangle)

	fmt.Printf("Triangle ((%f, %f), (%f, %f), (%f, %f)) area is: %f.", triangle.x1, triangle.y1, triangle.x2, triangle.y2, triangle.x3, triangle.y3, p)
}

func calculateArea(triangle Triangle) float64 {
	a := math.Sqrt(math.Pow(triangle.x1-triangle.x2, 2) + math.Pow(triangle.y1-triangle.y2, 2))
	b := math.Sqrt(math.Pow(triangle.x1-triangle.x3, 2) + math.Pow(triangle.y1-triangle.y3, 2))
	c := math.Sqrt(math.Pow(triangle.x2-triangle.x3, 2) + math.Pow(triangle.y2-triangle.y3, 2))

	s := (a + b + c) / 2

	p := math.Sqrt(s * (s - a) * (s - b) * (s - c))

	return p
}
