package main

import "fmt"

func main() {
	sayHello := func() {
		fmt.Println("hello")
	}
	sayHello()
	go sayHello()

	go func() {
		fmt.Println("hello")
	}()
}

func sayHello() {
	fmt.Println("hello")
}
