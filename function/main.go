package main

import "fmt"

func main() {

	fmt.Println("Welcome to the function in golang")
	gretter()
	result := proAdder(4, 5, 7, 8, 3)

	fmt.Println("The answer is ", result)

	result2, _ := proAdder2(4, 8, 7, 9, 10)
	fmt.Println("Another version of function", result2)

}

func added(valOne int, valTwo int) int { // her int outside() is the return type of function
	return (valOne + valTwo)
}

func gretter() {

	fmt.Println(" Hello Ankit welcome to golang")
}

// If i don't know how many value going to come  then we need to use Pro Function
func proAdder(values ...int) int {

	total := 0

	for i := 0; i < len(values); i++ {
		total += values[i]
	}

	return total
}

// if we have to return more then 1 values
func proAdder2(values ...int) (int, string) {

	total := 0

	for i := 0; i < len(values); i++ {
		total += values[i]
	}

	return total, "my name is ankit"
}
