package main

import "fmt"

func main() {
	i1 := 42
	s1 := []int{1, 2, 3, 4, 5, -5, 42}
	fmt.Println("Before changes:", i1)
	doublePtr(&i1)
	fmt.Println("After first call:", i1)
	i1 = doubleRet(i1)
	fmt.Println("After second change", i1)
	fmt.Println("Slice before change:", s1)
	doubleSlice(s1)
	fmt.Println("Slice after change:", s1)
}

func doublePtr(n *int) {
	*n *= 2
}

func doubleRet(n int) int {
	return n * 2
}

func doubleSlice(s []int) {
	for i := range s {
		s[i] *= 2
	}
}
