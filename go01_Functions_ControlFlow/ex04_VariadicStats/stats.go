package main

import "math"

func Mean(nums ...float64) float64 {
	sum := 0.0
	for _, num := range nums {
		sum += num
	}
	return sum / float64(len(nums))
}

func Variance(nums ...float64) float64 {
	avg := Mean(nums...)
	sum := 0.0
	for _, num := range nums {
		distance := num - avg
		sum += distance * distance
	}
	return sum / float64(len(nums))
}

func StdDev(nums ...float64) float64 {
	variance := Variance(nums...)
	return math.Sqrt(variance)
}
