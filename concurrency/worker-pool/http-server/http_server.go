package main

import (
	"fmt"
	"net/http"
	"time"
)

type Job struct {
	id      int
	request http.Request
}

type Result struct {
	jobID int
	data  string
}

func worker(workerId int, jobs chan Job, results chan Result) {
	for job := range jobs {
		result := process(workerId, job)
		results <- result
	}
}

func process(workerId int, job Job) Result {
	fmt.Printf("Processing job %d in worker %d\n", job.id, workerId)
	time.Sleep(1 * time.Second)
	return Result{job.id, "done"}
}

func main() {
	jobs := make(chan Job, 100)
	results := make(chan Result, 100)

	//start workers
	for w := 1; w <= 5; w++ {
		go worker(w, jobs, results)
	}
	fmt.Printf("---------Workers are ready--------\n")

	//Simulate incoming requests
	for i := 0; i < 10; i++ {
		fmt.Printf(">> sending request %d\n", i)
		jobs <- Job{i, http.Request{}}
	}

	close(jobs)

	for i := 0; i < 10; i++ {
		result := <-results
		fmt.Printf("<< Result for job %d: %s\n", result.jobID, result.data)
	}
}
