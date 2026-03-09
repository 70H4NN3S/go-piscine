package main

import "fmt"

func NewMatrix(rows, cols int) [][]float64 {
	m := make([][]float64, rows)
	for i := range m {
		m[i] = make([]float64, cols)

		// Fill with data to test
		for j := range m[i] {
			m[i][j] = float64((i*j + i + j) % 10)
		}
	}
	return m
}

func PrintMatrix(m [][]float64) {
	for range m[0] {
		fmt.Print("--")
	}
	fmt.Print("-\n")

	for i := range m {
		fmt.Println(m[i])
	}
}
