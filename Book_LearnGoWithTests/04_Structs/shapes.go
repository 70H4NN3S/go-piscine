// package shapes handles simple geographic shapes
package shapes

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rect Rectangle) float64 {
	width := rect.Width
	height := rect.Height
	return 2 * (width + height)
}

func Area(rect Rectangle) float64 {
	width := rect.Width
	height := rect.Height
	return width * height
}
