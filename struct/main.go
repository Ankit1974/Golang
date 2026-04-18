package main

import "fmt"

// In Go, %+v is a formatting verb used with fmt.Printf (or fmt.Sprintf) to print structs with their field names included.

/*
package main

import "fmt"

	type Person struct {
	    Name string
	    Age  int
	}

	func main() {
	    p := Person{Name: "Ankit", Age: 25}
	    fmt.Printf("%v\n", p)   // Output: {Ankit 25}
	    fmt.Printf("%+v\n", p)  // Output: {Name:Ankit Age:25}
	}
*/

func main() {
	fmt.Println("We are going to learn about struct in golang")

	// how user can use it

	ankit := User{"Ankit Raj", 33, "7D", "Maths"}

	fmt.Printf("Printing the struct %+v\n", ankit)

	fmt.Printf("my name is %v and roll number is %v", ankit.Name, ankit.Roll)

}

// createing struct

type User struct {
	Name    string
	Roll    int
	Class   string
	Subject string
}
