package main

import (
	"fmt"
)

func main() {
	a := []int{123, 1, 12, 43}
	histogram(a)
}

func histogram(bars []int) {
	maxN := getMax(bars)
	for index, bar := range bars {
		count := getCount(bar, maxN)
		printer(index, bar, count)
	}
}

func printer(index, value, count int) {
	fmt.Printf("Index %d | value=%d: ", index, value)
	for range count {
		fmt.Printf("#")
	}
	fmt.Printf("\n")
}

func getMax(bars []int) int {
	maxN := 0
	for _, n := range bars {
		if n > maxN {
			maxN = n
		}
		if -1*n > maxN {
			maxN = -1 * n
		}
	}
	return maxN
}

func getCount(bar, maxN int) int {
	if maxN == 0 {
		return 40
	}
	if bar < 0 {
		return -1 * bar * 40 / maxN
	}
	return bar * 40 / maxN
}
