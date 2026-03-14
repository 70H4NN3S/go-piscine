package main

import "fmt"

func main() {
	m1 := NewMatrix(5, 5)
	m2 := NewMatrix(6, 9)

	PrintMatrix(m1)
	PrintMatrix(Transpose(m1))
	multM1, err := MatMul(m1, Transpose(m1))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		PrintMatrix(multM1)
	}

	PrintMatrix(m2)
	PrintMatrix(Transpose(m2))
	multM2, err := MatMul(m2, Transpose(m2))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		PrintMatrix(multM2)
	}
}
