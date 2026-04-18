package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var number = []int{1, 2, 4, 5}

	fmt.Printf("Type of number is %T\n", number)

	number = append(number, 6, 7)

	number = number[1:3]
	fmt.Println(number)

	// Another way of creating slice
	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	highScores = append(highScores, 555, 666, 321)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)

	// How to remove value from slice based on the index
	var rollNumber = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(rollNumber)
	var index int = 4
	rollNumber = append(rollNumber[:index], rollNumber[index+1:]...)
	fmt.Println(rollNumber)
}
