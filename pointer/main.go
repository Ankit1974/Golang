package main

import "fmt"

// When we are referencing something, then we use & (address-of operator).
// when we have to get the actual vlaue of a pointer we *nameOfThePointer

func main() {

	fmt.Println(" Welcome to a class on Pointer ")

	// when no vlaue is assigned or alloted then it has nil value
	var ptr *int
	fmt.Println("value  Pointer is ", ptr)

	myNumber := 29
	var ptr2 = &myNumber
	fmt.Println(" address of actual pointer is", ptr2)

	*ptr2 = *ptr2 + 10

	fmt.Println("the new value is ", myNumber)

}
