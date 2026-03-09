package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return r.Width*2 + r.Height*2 }
func (r Rectangle) IsSquare() bool     { return r.Height == r.Width }

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64          { return c.Radius * c.Radius * math.Pi }
func (c Circle) Circumference() float64 { return 2 * c.Radius * math.Pi }

type Triangle struct {
	A float64
	B float64
	C float64
}

func (t Triangle) Area() float64 {
	s := (t.A + t.B + t.C) / 2
	square := s * (s - t.A) * (s - t.B) * (s - t.C)
	result := math.Sqrt(square)
	return result
}

func (t Triangle) IsValid() bool {
	switch {
	case t.A <= 0:
		return false
	case t.B <= 0:
		return false
	case t.C <= 0:
		return false
	case t.A+t.B <= t.C:
		return false
	case t.B+t.C <= t.A:
		return false
	case t.A+t.C <= t.B:
		return false
	default:
		return true
	}
}
