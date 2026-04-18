package main

import "fmt"

// how to declare map
func main() {
	nameOfTheStudent := make(map[int]string)

	nameOfTheStudent[1] = "Ankit Raj"
	nameOfTheStudent[2] = "Ankit Kumar"
	nameOfTheStudent[3] = "Ankit Sharama"

	// this is the way to acess the particular index in map

	fmt.Println("My name is ankit raj", nameOfTheStudent[2])

	//. this is the way to delete from particular index in map

	delete(nameOfTheStudent, 1)

	// this is the way to itterate over the map

	for _, value := range nameOfTheStudent {
		fmt.Println("this is my name", value)
	}
}
