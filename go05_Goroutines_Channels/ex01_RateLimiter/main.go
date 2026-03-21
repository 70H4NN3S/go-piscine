package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	mu       sync.Mutex
	cond     *sync.Cond
	capacity int
	maxCap   int
	refill   float64
}

func NewRateLimiter(capacity int, refillPerSec float64) *RateLimiter {
	r := &RateLimiter{
		capacity: capacity,
		maxCap:   capacity,
		refill:   refillPerSec,
	}
	r.cond = sync.NewCond(&r.mu)
	return r
}

func (r *RateLimiter) Allow() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.capacity > 0
}

func (r *RateLimiter) Wait() {
	r.mu.Lock()
	defer r.mu.Unlock()
	for r.capacity <= 0 {
		r.cond.Wait()
	}
}

func (r *RateLimiter) Increment() {
	for {
		time.Sleep(time.Second)
		r.mu.Lock()
		r.capacity += int(r.refill)
		if r.capacity > r.maxCap {
			r.capacity = r.maxCap
		}
		r.cond.Broadcast()
		r.mu.Unlock()
	}
}

func (r *RateLimiter) Request() {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.capacity <= 0 {
		fmt.Println("waiting...")
		r.cond.Wait()
	}
	r.capacity--
	fmt.Println("requesting....")
}

func main() {
	r := NewRateLimiter(10, 5)
	go r.Increment()
	for range 20 {
		r.Request()
	}
	time.Sleep(2 * time.Second)
	for range 15 {
		r.Request()
	}
}
