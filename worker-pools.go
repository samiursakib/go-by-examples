package main

import (
	"fmt"
	"time"
)

func worker_(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "job started", j)
		time.Sleep(1 * time.Second)
		fmt.Println("worker", id, "job finished", j)
		results <- j * 2
	}
}

func workerPools() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker_(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
