package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const url = "https://example.com"

func main() {

	fmt.Println(" learning web request ")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Printf("Response id a type of : %T\n", response)

	//databytes, err := io.ReadAll(response.Body) // if the data is small then use this

	databytes, err := io.Copy(os.Stdout, response.Body) // if thedata is too large then we use this , this will stream the data

	if err != nil {
		panic(err)
	}

	// content := string(databytes) // if we are using this io.ReadAll(response.Body) then we need to convert it into string

	content := (databytes) // if we are using this  io.Copy(os.Stdout, response.Body) then we don't need to convert it into string

	fmt.Println((content))
}
