package channel

/*
   🔹 What is a Channel in Go?
       Simple Definition
        A channel is a way for goroutines to communicate with each other.

        It is a typed communication pipe.

        Think of it like:
        A tube
        A message queue
        A data pipeline
        A communication bridge between goroutines

    Formal Definition

      Channel is a synchronization mechanism in Go used to send and receive data between goroutines safely.

    Why Channels Exist
      Because Go follows this philosophy:
      “Do not communicate by sharing memory; share memory by communicating.”
    Instead of many goroutines touching the same variable:
     👉 They pass data through channels.
*/


// How to Create Channel

// Basic Channel Creation

   ch := make(chan int)

/*   Types of Channel Creation
       There are two main kinds of channels:
         1. Unbuffered
         2. Buffered
*/

// Unbuffered Channel

ch := make(chan int)

// Buffered Channel

ch := make(chan int, 5) // 5 is the buffer size




// to Check if Channel is closed or not -  package main

import "fmt"

func main() {
    ch := make(chan int)

    go func() {
        ch <- 10
        close(ch)
    }()

    v, ok := <-ch
    fmt.Println(v, ok)   // 10 true

    v, ok = <-ch
    fmt.Println(v, ok)   // 0 false  (closed)
}
