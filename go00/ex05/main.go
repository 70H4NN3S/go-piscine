package main

import (
	"fmt"
	"os"
	"strconv"
)

func printError(err error) {
	fmt.Fprintf(os.Stderr, err.Error()+"\n")
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		printError(fmt.Errorf("usage: ./ex05 <n>"))
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		printError(err)
	}

	current, previous := 1, 0
	if n < 1 {
		return
	}
	fmt.Println(previous)
	if n < 2 {
		return
	}
	for i := 0; i < (n - 1); i++ {
		fmt.Println(current)
		current, previous = current+previous, current
	}
}
