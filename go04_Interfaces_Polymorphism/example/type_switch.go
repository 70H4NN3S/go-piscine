package main

import "fmt"

func describe(i any) {
	// Safe assertion (comma-ok)
	if s, ok := i.(string); ok {
		fmt.Println("string:", s)
		return
	}
	switch v := i.(type) {
	case int:
		fmt.Printf("int: %d\n", v)
	case string:
		fmt.Printf("string: %q\n", v)
	case bool:
		fmt.Printf("bool: %t\n", v)
	default:
		fmt.Printf("unknown: %T\n", v)
	}
}

func main() {
	describe(nil)
	describe("Hello")
	describe(3.14)
	describe(true)
	describe(345)
}
