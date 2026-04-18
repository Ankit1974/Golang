package context

import (
	"context"
	"fmt"
	"time"
)

/*

  🔹 What is Context Package in Go?

  Simple Definition

  The context package in Go is used to:

    . Cancel long running operations
    . Pass deadlines
    . Pass request-scoped values
    . Coordinate multiple goroutines
    . Why Do We Need Context?

  Imagine:

  . You start 10 goroutines
  . User cancels request
  . Or timeout happens

❌ Without context – goroutines keep running forever

✔ With context – we can STOP them safely

Main Purpose:-

. Context gives a standard way to:
. Cancel work
. Set timeouts
. Propagate cancellation signals

*/

// Think this is a worker (chef)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Order cancelled – stop work")
			return

		default:
			fmt.Println("Cooking food...")
			time.Sleep(1 * time.Second)
		}
	}
}

// Main program (customer)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(5 * time.Second)

	cancel() // this triggers <-ctx.Done()

}

/*  Root:
- Background()
- TODO()

Incoming:
- Request.Context()

Derived:
- WithCancel → manual stop
- WithTimeout → auto stop
- WithDeadline → fixed time
- WithValue → pass data

*/

/* Context in Go is used for cancellation, timeouts, and passing request-scoped data.
We usually start with a root context like context.Background(),
but in APIs we use c.Request.Context() so cancellation propagates from the client.
Then we derive contexts using WithTimeout or WithCancel to control execution. */
