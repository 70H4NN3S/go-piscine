package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		printErr(fmt.Errorf("usage: ./ex00 <num1> <op> <num2>"))
	}
	num1, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		printErr(fmt.Errorf("<num1> has to be a valid number"))
	}
	num2, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		printErr(fmt.Errorf("<num2> has to be a valid number"))
	}
	switch os.Args[2] {
	case "+":
		addNums(num1, num2)
	case "-":
		subNums(num1, num2)
	case "x":
		mulNums(num1, num2)
	case "/":
		divNums(num1, num2)
	default:
		printErr(fmt.Errorf("<op> has to be a valid operator.\nSupported: (+,-,x,/)"))
	}
}

func addNums(num1, num2 float64) {
	result := num1 + num2
	printRes(result)
}

func subNums(num1, num2 float64) {
	result := num1 - num2
	printRes(result)
}

func mulNums(num1, num2 float64) {
	result := num1 * num2
	printRes(result)
}

func divNums(num1, num2 float64) {
	if num2 == 0 {
		printErr(fmt.Errorf("division by 0 impossible!"))
	}
	result := num1 / num2
	printRes(result)
}

func printErr(err error) {
	fmt.Fprintf(os.Stderr, err.Error()+"\n")
	os.Exit(0)
}

func printRes(result float64) {
	fmt.Printf("%0.2f\n", result)
}
