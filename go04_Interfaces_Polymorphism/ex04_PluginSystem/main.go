package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := NewRegistry()
	r.Register(Echo{})
	r.Register(Upper{})
	r.Register(Reverse{})

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		err := r.Execute(input)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
		}
	}
}
