package main

import "fmt"

// %T is used to get the type
// \n is used for next line

const LongString string = " Ankit Raj" // This Variable is public becuase  starting letter of the variable is capital

func main() {

	// String
	var name string = "Ankit Raj"
	fmt.Printf("this variable is a type of : %T \n", name)

	// bool
	var isLoggedIn bool = false
	fmt.Printf("this variable is a type of :%T \n", isLoggedIn)

	// smallFloat
	var smallFloat float32 = 233.845843
	fmt.Printf("this variable is a type of : %T \n", smallFloat)

	/* Default (zero) values in Go

	   1. int	0
	   2. string	"" (empty string)
	   3. bool	false
	   4. float	0.0
	   5. pointer	nil
	   6. slice	nil
	   7. map	nil

	*/

	// Default value of Integer ( if we don't allocate or declare any value)
	var defaultValue int
	fmt.Println(defaultValue)
	fmt.Printf(" this variable is a type of : %T \n", defaultValue)

	// Default value of String
	var defaultValue2 string
	fmt.Println(defaultValue2)
	fmt.Printf(" this variablr is a type of :  : %T \n", defaultValue2)

	// Another way of Declearing Variable ( implicit Decleration)

	var name2 = "Ankit Raj"
	fmt.Println(name2)

	// Another way of Declaring Variable ( Short Variable Declaration )
	// Notes - this is only possible under the function , we can't  declare  variable in this way outside the function
	name3 := "Ankit Raj"
	fmt.Println(name3)

}

/* In Go, capitalization controls visibility.

### Rule (super important in Go)

1. Capital letter → Exported (public)
2. Small letter → Unexported (private)

 This applies to:

1. Variables
2. functions
3. structs
4. struct fields
5. interfaces
6. constants

*/


