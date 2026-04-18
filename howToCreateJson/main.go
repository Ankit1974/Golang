package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type student struct {
	Name  string   `json:"studentName"` // this field name studentName should shown when some one use my api
	Roll  int      `json:"studentRollNumber"`
	Class int      `json:"class"`
	Email string   `json:"-"`              // when we want if anyone use my api they will not able to see this field
	Tags  []string `json:"tags,omitempty"` // If this field is empty, don’t include it in the JSON output.
}

func main() {
	fmt.Println(" we are going to learn how to create Json")
	EncodeJson()
	DecodeJsonFromAPI("https://api.example.com/student/1")

	// Dynamic key-value JSON examples
	EncodeKeyValueJson()
DecodeKeyValueJsonFromAPI("https://api.example.com/profile")
}

func EncodeJson() {

	studentList := []student{
		{"Ankit1", 32, 7, "Ankit@19741", []string{"good", "veryGood"}},
		{"Ankit2", 33, 8, "Ankit@19742", []string{"good", "averageGood"}},
		{"Ankit3", 34, 5, "Ankit@19743", []string{"bad"}},
		{"Ankit4", 35, 9, "Ankit@19745", []string{"goods", "veryGoods"}},
		{"Ankit5", 36, 10, "Ankit@19746", nil},
	}

	// Package this data as json data
	finaljson, err := json.MarshalIndent(studentList, "", "\t")
	if err != nil {
		log.Println("failed to encode json:", err)
		return
	}

	fmt.Printf("%s\n", finaljson)
}

func DecodeJsonFromAPI(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("failed to fetch api:", err)
		return
	}
	defer resp.Body.Close()

	var studentData student

	decoder := json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&studentData); err != nil {
		log.Println("failed to decode json:", err)
		return
	}

	fmt.Printf("Decoded from API: %#v\n", studentData)
}

// Some  cases where you just want to add data key value pair

func EncodeKeyValueJson() {
	// Some cases where you just want to add data key value pair
	// For example, dynamic payloads or unknown fields
	data := map[string]interface{}{
		"studentName":       "Ankit",
		"studentRollNumber": 32,
		"class":             7,
		"tags":              []string{"good", "veryGood"},
		"isActive":          true, // dynamically added field
	}

	// Convert map to JSON
	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Println("failed to encode json:", err)
		return
	}

	fmt.Println("Dynamic JSON payload:")
	fmt.Println(string(jsonData))
}

func DecodeKeyValueJsonFromAPI(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("failed to fetch api:", err)
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}

	decoder := json.NewDecoder(resp.Body)

	if err := decoder.Decode(&result); err != nil {
		log.Println("failed to decode json:", err)
		return
	}

	fmt.Println("Decoded key-value JSON from API:")
	for k, v := range result {
		fmt.Printf("%s: %#v\n", k, v)
	}
}
