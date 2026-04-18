package mutex

/* A Mutex in Go (short for Mutual Exclusion) is a synchronization mechanism used to protect shared data from being accessed by multiple goroutines at the same time.

   It is provided by the sync package as:

   var mu sync.Mutex

*/

/*
  🔴 The Problem Mutex Solves
Go programs often run many goroutines concurrently.
If two or more goroutines try to modify the same variable at the same time, it creates a race condition.
Example of a race condition:
var counter int

func increment() {
    counter = counter + 1
}

If 100 goroutines call increment() at the same time, the final value of counter may NOT be 100 because:
Reads and writes overlap Updates get lost.

*/

// When we perform any action on the database then First Lock it and after performing the action unlock it
