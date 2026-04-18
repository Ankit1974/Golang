package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println(" we are going study Switch Statment")
	rand.Seed(time.Now().UnixNano()) // this line will help to generate the random number.  Seed the random number generato , time.Now Current time , UnixNano time in nanoseconds
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dic is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("dice value is 1")
	case 2:
		fmt.Println("dice value is 2")
	case 3:
		fmt.Println("dicen value is 3")
		fallthrough // It forces execution to continue to the next case, even if that case’s condition doesn’t match.
	case 4:
		fmt.Println("dicen value is 4")
	case 5:
		fmt.Println("dicen value is 5")
	case 6:
		fmt.Println("dicen value is 6")
	default:
		fmt.Println("what was that!")

	}

}
