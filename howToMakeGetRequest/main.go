package main

import (
	"fmt"
	"io"
	"net/http"
)

func PerformGetRequest() {

	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Status:", response.Status)

	// Read response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response Body:")
	fmt.Println(string(body))
}
