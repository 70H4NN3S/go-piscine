package main

import "fmt"

type Animal struct {
	Name string
	Age  int
}

func (a Animal) Describe() string {
	return fmt.Sprintf("%s (age %d)", a.Name, a.Age)
}

type Dog struct {
	Animal // embedded - promotes fields & methods
	Breed  string
}

func main() {
	d := Dog{Animal: Animal{"Rex", 3}, Breed: "Labrador"}
	fmt.Println(d.Name)       // promoted from Animal
	fmt.Println(d.Describe()) // promoted method
}
