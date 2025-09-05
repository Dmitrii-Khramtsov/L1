package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

const nJobs = 10

func worker(id int, jobs chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d start job %d", id, job)
		fmt.Println(job)
		fmt.Printf("Worker %d finished job %d", id, job)
	}
}

func main() {
	jobs := make(chan int)

	wg := sync.WaitGroup{}
	for i := range nJobs {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	for range 10 {
		jobs <- rand.IntN(100)
	}
	close(jobs)

	wg.Wait()
}
