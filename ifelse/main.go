package main

import "fmt"

func main() {

	fmt.Println("Lets start ifelse in golang")

	loginCount := 5
	var result string

	if loginCount > 10 {
		result = " regular user"
	} else if loginCount < 6 {
		result = " less regular user"
	} else {
		result = " not a regular user"
	}

	fmt.Println(result)

	// Another way of doing if and else

	if num := 3; num < 10 {
		fmt.Println(" number is less then 10 ")
	} else {
		fmt.Println(" number if grater then 10")
	}

}
