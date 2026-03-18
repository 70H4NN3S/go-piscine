package main

import (
	"fmt"
	"time"
)

func example() {
	ch := make(chan int)        // unbuffered - synchronous
	bch := make(chan string, 3) // buffered - async up to 3

	// Range over a channel (close signals done)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	for v := range ch {
		fmt.Println(v)
	}

	// Select - multiplex channels
	select {
	case msg := <-ch1:
		fmt.Println("ch1:", msg)
	case msg := <-ch2:
		fmt.Println("ch2:", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	default:
		fmt.Println("no channel ready")
	}
}
