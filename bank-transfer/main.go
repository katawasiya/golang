package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func requestToAcquiring() <-chan int {
	responseChan := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		responseChan <- 1
	}()
	return responseChan
}

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		start := time.Now()
		<-requestToAcquiring()
		elapsed := time.Since(start)
		fmt.Printf("Paymant %d finished bank transfer %d in %v\n", id, j, elapsed)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	const numJobs = 1000
	jobs := make(chan int, numJobs)

	var wg sync.WaitGroup
	wg.Add(2)

	for w := 1; w <= 2; w++ {
		go worker(w, jobs, &wg)
	}

	go func() {
		for j := 1; j <= numJobs; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	wg.Wait()
}
