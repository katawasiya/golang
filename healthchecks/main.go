package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"pornoHub.com",
		"XVideos.com",
		"Porno365.com",
		"Redtube.com",
	}

	statuses := parallelCheck(urls)

	for status := range statuses {
		fmt.Println(status)
	}
}

func parallelCheck(urls []string) <-chan int {
	statusChannel := make(chan int)
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()
			statusChannel <- healthCheck(url)
		}(url)
	}

	go func() {
		wg.Wait()
		close(statusChannel)
	}()

	return statusChannel
}

func healthCheck(url string) int {
	time.Sleep(2 * time.Second)
	return http.StatusCreated
}
