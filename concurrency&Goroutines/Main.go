// without mutex 
package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Result represents the response from an API call
type Result struct {
	URL      string
	Duration time.Duration
	Status   int
	Err      error
}

func main() {
	endpoints := []string{
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://jsonplaceholder.typicode.com/users/1",
		"https://jsonplaceholder.typicode.com/comments/1",
		"https://jsonplaceholder.typicode.com/albums/1",
		"https://jsonplaceholder.typicode.com/todos/1",
		"https://jsonplaceholder.typicode.com/photos/1",
	}

	fmt.Println("--- Starting Sequential API Calls ---")
	start := time.Now()
	for _, url := range endpoints {
		fetchSync(url)
	}
	fmt.Printf("Sequential total time: %v\n\n", time.Since(start))

	fmt.Println("--- Starting Concurrent API Calls (Goroutines) ---")
	start = time.Now()
	results := fetchAsync(endpoints)

	for res := range results {
		if res.Err != nil {
			fmt.Printf("FAIL: %s | Error: %v\n", res.URL, res.Err)
		} else {
			fmt.Printf("SUCCESS: %s | Status: %d | Latency: %v\n", res.URL, res.Status, res.Duration)
		}
	}
	fmt.Printf("Concurrent total time: %v\n", time.Since(start))
}

// fetchSync performs a sequential (blocking) API call
func fetchSync(url string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("Fetched %s | Status: %d | Time: %v\n", url, resp.StatusCode, time.Since(start))
}

// fetchAsync performs concurrent API calls using Goroutines and WaitGroups
func fetchAsync(urls []string) <-chan Result { // <-chan Result means it will return a channel of Result type, only reading from the channel
	resultChan := make(chan Result, len(urls)) // Buffer channel to prevent blocking , means it can store reuslt upto the length of url slice after that it will block
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		// Launching each API call in its own Goroutine
		go func(targetURL string) {
			defer wg.Done()

			start := time.Now()
			resp, err := http.Get(targetURL)

			res := Result{URL: targetURL, Duration: time.Since(start)}
			if err != nil {
				res.Err = err
			} else {
				res.Status = resp.StatusCode
				resp.Body.Close()
			}

			resultChan <- res
		}(url)
	}

	// Close the channel once all Goroutines are finished
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	return resultChan
}
