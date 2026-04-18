package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")

	var fruitlist [4]string

	fruitlist[0] = "mango"
	fruitlist[1] = "orange"
	fruitlist[3] = "Peach"

	fmt.Println("Fruits list is", fruitlist)
	fmt.Println("Fruits count is", len(fruitlist)) // it always give number that we have assign suring array inslization

	// Initialize and declare the array at the same time (compiler infers length)
	var number = [...]int{1, 2, 3, 4}
	fmt.Println("number list is", number)
}
