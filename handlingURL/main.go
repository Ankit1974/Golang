// package main

// import (
// 	"fmt"
// 	"net/url"
// )

// const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

// func main() {

// 	fmt.Println(" We are going to learn how to handle url")
// 	fmt.Println(myurl)

// 	//parsing
// 	result, _ := url.Parse(myurl)

// 	// if err != nil {

// 	// 	panic(err)
// 	// }

// 	fmt.Println(result.Scheme)
// 	fmt.Println(result.Host)
// 	fmt.Println(result.Path)
// 	fmt.Println(result.Port())
// 	fmt.Println(result.RawQuery)
// }

package main

import (
	"fmt"
	"net/url"
)

func main() {

	fmt.Println("We are going to learn how to handle url")

	myURL := "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

	//Parsing
	result, err := url.Parse(myURL)
	if err != nil {
		panic(err) // or handle gracefully
	}

	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Path:", result.Path)
	fmt.Println("Raw Query:", result.RawQuery)

	// Query params
	queryParams := result.Query()
	fmt.Println("Course name:", queryParams.Get("coursename"))
	fmt.Println("Payment ID:", queryParams.Get("paymentid"))

	// Applying for loop

	for _, value := range queryParams {
		fmt.Println("paramas is ", value)
	}

	// if we have all the infromation and we want to create URL

	partsOfUrk := &url.URL{

		Scheme:  "https",
		Host:    "ankit.dev",
		Path:    "/goe",
		RawPath: "user=Ankit",
	}

	// Query parameters
	q := partsOfUrk.Query()
	q.Set("coursename", "reactjs")
	q.Set("paymentid", "ghbj456ghb")

	partsOfUrk.RawQuery = q.Encode()

	creatingUrl := partsOfUrk.String()
	fmt.Println(creatingUrl)
}
