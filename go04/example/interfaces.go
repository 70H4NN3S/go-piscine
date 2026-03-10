package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

// Both types implicitly implement Shape
type (
	Rect   struct{ W, H float64 }
	Circle struct{ R float64 }
)

func (r Rect) Area() float64        { return r.W * r.H }
func (r Rect) Perimeter() float64   { return 2 * (r.W * r.H) }
func (c Circle) Area() float64      { return math.Pi * c.R * c.R }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.R }

func printShape(s Shape) {
	fmt.Printf("Area: %.2f Perim: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	r := Rect{6.4, 3.14}
	c := Circle{3.45}
	printShape(r)
	printShape(c)
}
