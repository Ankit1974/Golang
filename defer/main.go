package main

/*
    Why defer is AWESOME 🔥

✅ Cleanup always happens
✅ Even if an error occurs
✅ Even if function returns early

*/

/*

   Notes :-

   Multiple defer calls run in reverse order (stack) - LIFO.

*/

// Defination - “Do this later, just before the function finishes.”

import (
	"fmt"
)

func main() {
	defer fmt.Println("World")
	fmt.Println("hello")
}
