package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) Distance(other *Point) float64 {
	dx := other.x - p.x
	dy := other.y - p.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	a := NewPoint(1, 2)
	b := NewPoint(4, 6)

	fmt.Println(a.Distance(b))
}
