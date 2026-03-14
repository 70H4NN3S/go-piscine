package main

import "fmt"

func main() {
	hasCar := true
	age := 21
	height := 175.3
	name := "Julian"
	var cheap byte = 'A'
	favSymbol := 'Ä'

	fmt.Printf("bool: %v (type: %T)\n", hasCar, hasCar)
	fmt.Printf("int: %v (type: %T)\n", age, age)
	fmt.Printf("float64: %v (type: %T)\n", height, height)
	fmt.Printf("string: %v (type: %T)\n", name, name)
	fmt.Printf("byte: %v (type: %T)\n", cheap, cheap)
	fmt.Printf("rune: %v (type: %T)\n", favSymbol, favSymbol)
}
