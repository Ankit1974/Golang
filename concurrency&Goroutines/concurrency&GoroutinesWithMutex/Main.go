// with mutex
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
		"https://jsonplaceholder.typicode.com/posts/2",
		"https://jsonplaceholder.typicode.com/users/2",
		"https://jsonplaceholder.typicode.com/comments/2",
		"https://jsonplaceholder.typicode.com/albums/2",
		"https://jsonplaceholder.typicode.com/todos/2",
		"https://jsonplaceholder.typicode.com/photos/2",
	}

	// Shared resources
	var allResults []Result
	var mu sync.Mutex // The Mutex protects 'allResults'
	var wg sync.WaitGroup

	fmt.Println("--- Starting Concurrent API Calls (Mutex) ---")
	start := time.Now()

	for _, url := range endpoints {
		wg.Add(1)
		go func(targetURL string) {
			defer wg.Done()

			// 1. Perform the external work (API call) outside the lock
			fetchStart := time.Now()
			resp, err := http.Get(targetURL)

			res := Result{
				URL:      targetURL,
				Duration: time.Since(fetchStart),
			}

			if err != nil {
				res.Err = err
			} else {
				res.Status = resp.StatusCode
				resp.Body.Close()
			}

			// 2. Lock the mutex before writing to the shared slice
			mu.Lock()
			allResults = append(allResults, res)
			mu.Unlock() // 3. Important: Unlock as soon as the write is done
		}(url)
	}

	// Wait for all Goroutines to finish
	wg.Wait()

	totalTime := time.Since(start)

	// Now we can safely read the results
	for _, res := range allResults {
		if res.Err != nil {
			fmt.Printf("FAIL: %s | Error: %v\n", res.URL, res.Err)
		} else {
			fmt.Printf("SUCCESS: %s | Status: %d | Latency: %v\n", res.URL, res.Status, res.Duration)
		}
	}

	fmt.Printf("\nTotal time with Mutex: %v\n", totalTime)
}
