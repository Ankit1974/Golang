package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to user Input "
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Rating of the Book")

	// commma of ||  comma err

	input, _ := reader.ReadString('\n')
	fmt.Println(" Thanks for Rating ", input)

}


