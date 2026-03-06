package main

import (
	"fmt"
	"os"
	"strconv"
)

func errorPrint(s string) {
	fmt.Fprintln(os.Stderr, s)
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		errorPrint("usage: ./ex04 <celsiu>\ngo run main.go <celsius>")
	}

	celsius, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		errorPrint("Unable to parse Number. Type in any floating number")
	}
	fahrenheit := (celsius * 9 / 5) + 32
	fmt.Printf("%0.2f° Celcius is %0.2f° Fahrenheit\n", celsius, fahrenheit)
}
