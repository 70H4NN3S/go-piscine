package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func submitSquareRoot(input int, pool *Pool) {
	for i := range input {
		payload := rand.Float64() * 10000
		job := Job{i, payload}
		pool.Submit(job)
	}
}

func main() {
	squareRoot := func(j Job) Result {
		n, ok := j.Payload.(float64)
		if !ok {
			return Result{j.ID, 0, fmt.Errorf("the payload is not of type float64")}
		}
		res := math.Sqrt(n)
		return Result{j.ID, res, nil}
	}
	pool := NewPool(5, squareRoot)

	fmt.Println("### WORKER POOL STARTED ###")
	t := time.Now()
	go func() {
		submitSquareRoot(50, pool)
		pool.Stop()
	}()

	fmt.Println("### WORKER POOL STARTED ###")
	for r := range pool.Results() {
		if r.Err != nil {
			fmt.Printf("Something went wrong wit ID %d ==> %s\n", r.JobID, r.Err.Error())
		} else {
			fmt.Printf("ID: %d; Output: %.2f\n", r.JobID, r.Output)
		}
	}
	fmt.Println("### WORKER POOL FINISHED IN", time.Since(t), "###")
}
