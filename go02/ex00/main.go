package main

import "fmt"

func main() {
	arr1 := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		slice = Insert(slice, i, i*i)
	}
	fmt.Println(arr1)
	fmt.Println(Remove(arr1[:], 3))
	fmt.Println("-----------------\n\n\n\n")
	fmt.Println(Unique([]int{1, 2, 3, 1, 2, 1, 2, 3, 1, 2, 3, 4, 2, 1, 5, 2, 1}))
	fmt.Println("-----------------\n\n\n\n")
	fmt.Println(slice)
	fmt.Println("_________________\n\n\n\n")
	fmt.Println(Reverse(slice))
}

func Remove(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}

func Insert(s []int, i int, val int) []int {
	value := []int{val}
	second := append(value, s[i:]...)
	return append(s[:i], second...)
}

func Unique(s []int) []int {
	res := []int{}
	m := make(map[int]bool)
	for _, i := range s {
		if _, ok := m[i]; !ok {
			m[i] = true
			res = append(res, i)
		}
	}
	return res
}

func Reverse(s []int) []int {
	last := len(s) - 1
	for i := 0; i <= len(s)/2; i++ {
		s[i], s[last-i] = s[last-i], s[i]
	}
	return s
}
