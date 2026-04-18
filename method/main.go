package main

import "fmt"

func main() {
	fmt.Println("We are going to learn about struct in golang")

	// how user can use it

	ankit := User{"Ankit Raj", 33, "7D", "Maths"}

	fmt.Printf("Printing the struct %+v\n", ankit)

	fmt.Printf("my name is %v and roll number is %v\n", ankit.Name, ankit.Roll)

	ankit.GetClass()
	ankit.NewRoll()

}

// createing struct

type User struct {
	Name    string
	Roll    int
	Class   string
	Subject string
}

func (v User) GetClass() {

	fmt.Println("Class of the User", v.Class)
}

// Doing something intresting and expermintal

func (t User) NewRoll() {

	t.Roll = 23 // This Are creating copy not changing the actual value
	fmt.Println(" my new rollNumber is ", t.Roll)

}
