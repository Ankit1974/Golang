package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println(" we are going to learn file in golang")
	content := ("My Name is Ankit Raj , I am  Learning golang to understand the backend devlopment ")

	file, err := os.Create("./ankit.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("length is: ", length)

	defer file.Close()

	readfile("./ankit.txt")

}

func readfile(filname string) {

	databyte, err := os.ReadFile(filname)

	checkNilErr(err)

	fmt.Println("Text data inside the file \n", string(databyte))

}

func checkNilErr(err error) {

	if err != nil {

		panic(err)
	}
}
