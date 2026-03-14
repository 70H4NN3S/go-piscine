package main

import "fmt"

func main() {
	data1 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, -1}
	data2 := []float64{5, 4, 5, 6, 5, 4, 5, 6, 5, 4, 5, 6}
	data3 := []float64{1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9}

	printData(data1)
	printData(data2)
	printData(data3)
}

func printData(nums []float64) {
	mean := Mean(nums...)
	variance := Variance(nums...)
	stdDev := StdDev(nums...)

	fmt.Println(nums)
	fmt.Printf("%-10s %-10s %-10s\n", "Mean", "variance", "StdDev")
	fmt.Printf("%-10.2f %-10.2f %-10.2f\n", mean, variance, stdDev)
	fmt.Print("----------------------------------------\n\n\n")
}
