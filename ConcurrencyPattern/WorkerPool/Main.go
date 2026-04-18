package workerpool

/*
   🔹 What is a Worker Pool?

         Simple Definition

         A worker pool is a group of fixed number of goroutines that repeatedly take tasks from a queue (channel) and process them.
         Instead of creating unlimited goroutines, we create:

       👉 A limited set of workers that reuse themselves.

       Why Do We Need Worker Pool?

      Imagine you have:
        10,000 tasks
        If you create 10,000 goroutines → ❌ too heavy
        System may slow down or crash

        Better idea:

      👉 Create only 10 or 20 workers and let them handle all tasks.

      Real World Analogy

      Think of:

        A restaurant kitchen
        3 chefs (workers)
        100 orders (tasks)
        Orders come in a queue
    Chefs pick one by one and finish them.


*/

/*
    Components of Worker Pool :-

   1. A worker pool normally has:

   2. Job queue (channel)

   3. Fixed number of workers

   4. Result collection

   5. WaitGroup
*/

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
	}
}

func main() {
	const numWorkers = 3

	jobs := make(chan int, 10)
	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Send jobs
	for j := 1; j <= 10; j++ {
		jobs <- j
	}

	close(jobs)

	wg.Wait()
}
