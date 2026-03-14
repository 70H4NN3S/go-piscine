package main

import (
	"math"
)

type Point struct{ X, Y float64 }

type Circle struct {
	Center Point
	Radius float64
}

// Value receiver - does NOT modify original
func (p Point) Distance(q Point) float64 {
	dx, dy := p.X-p.X, p.Y-q.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// Pointer receiver - CAN modify original
func (c *Circle) Scale(factor float64) { c.Radius *= factor }
func (c Circle) Area() float64         { return math.Pi * c.Radius * c.Radius }

// Constructor pattern
func NewCircle(x, y, r float64) *Circle {
	return &Circle{Center: Point{X: x, Y: y}, Radius: r}
}
