package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func UploadMultipartData() {

	/*
	   . This is an in-memory buffer
	   . HTTP request body must be an io.Reader
	   . bytes.Buffer implements io.Reader
	*/
	body := &bytes.Buffer{}

	/*
		         . Manages multipart boundaries
		         . Encodes data in multipart/form-data format
		         . Writes directly into body

				 Without this → ❌ browser/server won’t understand files
	*/
	writer := multipart.NewWriter(body)

	// -------- TEXT FIELDS --------
	writer.WriteField("username", "ankit")
	writer.WriteField("email", "ankit@gmail.com")

	// -------- MULTIPLE PHOTOS --------
	photos := []string{"photo1.jpg", "photo2.jpg"}

	for _, photoPath := range photos {
		file, err := os.Open(photoPath) // Opens file from disk
		if err != nil {
			panic(err)
		}
		defer file.Close()

		/*
		   Creates a new multipart section
		   Field name = "photos"
		   Filename = photoPath
		*/
		part, err := writer.CreateFormFile("photos", photoPath)
		if err != nil {
			panic(err)
		}
		/*
		   Streams file data into request body
		   No full file load into memory
		   Efficient and safe
		*/
		io.Copy(part, file)
	}

	// -------- DOCUMENT --------
	doc, err := os.Open("resume.pdf")
	if err != nil {
		panic(err)
	}
	defer doc.Close()

	docPart, err := writer.CreateFormFile("document", "resume.pdf")
	if err != nil {
		panic(err)
	}
	io.Copy(docPart, doc)

	// close multipart writer
	writer.Close()

	// -------- HTTP REQUEST --------
	/*
	   Method: POST
	   URL: /upload
	   Body: multipart data
	*/
	req, err := http.NewRequest(
		"POST",
		"http://localhost:8000/upload",
		body,
	)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Upload Status:", resp.Status)
}

func main() {
	UploadMultipartData()
}
