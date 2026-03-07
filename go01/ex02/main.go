package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 1 && len(os.Args) != 3 {
		printErr(fmt.Errorf("usage:\n./ex02\n./ex02 <base> <exp>\n"))
	}
	if len(os.Args) == 3 {
		num1, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			printErr(fmt.Errorf("Failed to convert <base>. Use int\n"))
		}
		num2, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			printErr(fmt.Errorf("Failed to convert <exp>. Use int}n"))
		}
		fmt.Println(recursivePOW(num1, num2))
	} else {
		for i := 0; i <= 10; i++ {
			res := 1
			for j := 0; j < i; j++ {
				res *= 2
			}
			fmt.Println(res)
		}
	}
}

func recursivePOW(value, exp int64) int64 {
	if exp < 1 {
		return 1
	}
	if exp%2 == 0 {
		half := recursivePOW(value, exp/2)
		return half * half
	}
	return value * recursivePOW(value, exp-1)
}

func printErr(err error) {
	fmt.Fprintf(os.Stderr, err.Error())
	os.Exit(0)
}
