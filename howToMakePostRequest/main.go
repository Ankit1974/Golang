package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Course struct {
	Name   string `json:"name"`
	Course string `json:"course"`
	Price  int    `json:"price"`
}

func PerformPostRequest() {

	url := "http://localhost:8000/post"

	//  Create payload (struct → JSON)
	payload := Course{
		Name:   "Ankit",
		Course: "Go Backend",
		Price:  4999,
	}

	jsonData, err := json.Marshal(payload) // Marshal convert Go data → JSON
	if err != nil {
		panic(err)
	}

	//  Send POST request
	response, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(jsonData), //  this will turn JSON data into something HTTP can read from
	)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	//  Read response
	fmt.Println("Status Code:", response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response:", string(body))
}

func main() {
	PerformPostRequest()
}
