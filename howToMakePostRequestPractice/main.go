package howtomakepostrequestpractice

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Course struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func PostRequest() {

	myurl := "kjjkfe..//..//.//kjejr"

	payload := Course{
		Name:  "javascript",
		Price: 344,
	}

	jsonData, err := json.Marshal(payload)

	if err != nil {
		print(err)
	}

	response, err := http.Post(
		myurl,
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

}
