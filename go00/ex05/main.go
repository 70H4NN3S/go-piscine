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
		printError(fmt.Errorf("Invalid amount of arguments. Required args: 2"))
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		printError(err)
	}
	if n < 2 {
		return
	}

	current, previous := 1, 0
	fmt.Println(previous)
	for i := 0; i < (n - 1); i++ {
		fmt.Println(current)
		current, previous = current+previous, current
	}
}
