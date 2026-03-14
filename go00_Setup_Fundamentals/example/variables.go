package main

func main() {
	// Explicit declaration
	var name string = "Go"
	var age int = 15

	// Short variable declaration (:=)
	language := "Go"
	year := 2009

	// Multiple assignment
	x, y := 10, 20

	// Constants
	const Pi = 3.14159
	const Hello = "Hello, World"

	// Zero values: int=0, string="", bool=false
	var i int
	var s string
	var b bool
	_ = i
	_ = s
	_ = b
	_ = name
	_ = age
	_ = language
	_ = year
	_ = x
	_ = y
}
