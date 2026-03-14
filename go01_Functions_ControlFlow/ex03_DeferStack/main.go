package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}

	cleanup("Johannes")
	cleanup("Gabriel")
	cleanup("Guillermo")
}

func cleanup(name string) {
	defer fmt.Printf("Cleaning up: %s\n", name)
}
