package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	sayHello1 := func() {
		defer wg.Done()
		fmt.Println("hello")
	}
	wg.Add(1)
	go sayHello1()
	go sayHello2(&wg)
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("hello")
	}()
	wg.Add(1)
	wg.Wait()
}

func sayHello2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hello")
}
