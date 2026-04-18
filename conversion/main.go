package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println(" Rate my Pizza App ")

	Reader := bufio.NewReader(os.Stdin)

	fmt.Println(" Enter the rating ")

	input, _ := Reader.ReadString('\n')
	fmt.Println("Thnaks for Rating ", input)

	// Converting String into number

	numberRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Add one to my Rating", numberRating+1)
	}

}
