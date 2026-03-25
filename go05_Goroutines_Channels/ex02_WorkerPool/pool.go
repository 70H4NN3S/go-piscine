package main

import (
	"fmt"
	"sync"
)

type Job struct {
	ID      int
	Payload any
}

type Result struct {
	JobID  int
	Output any
	Err    error
}

type Pool struct {
	Input  chan Job
	Result chan Result
	wg     sync.WaitGroup
}

func NewPool(workers int) *Pool {
	input := make(chan Job, 5)
	result := make(chan Result, 5)
	p := &Pool{input, result, sync.WaitGroup{}}

	for range workers {
		p.wg.Add(1)
		go p.Worker()
	}
	return p
}

func (p *Pool) Worker() {
	defer p.wg.Done()
	for job := range p.Input {
		fmt.Println(job)
	}
}

func (p *Pool) Submit(job Job) {
}

func Results() <-chan Result {
	return nil
}

func (p *Pool) Stop() {
	close(p.Input)
	p.wg.Wait()
	close(p.Result)
}
