package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(" Welcme to the time study of golang")

	// This Will give the pargentTime
	pargentTime := time.Now()

	fmt.Println(pargentTime)

	// This Will Helps to Fomate the Time
	fmt.Println(pargentTime.Format("01-02-2006 Monday"))

	// Thsi will helps to create Date
	createDate := time.Date(2025, time.March, 12, 23, 0, 0, 0, time.UTC)

	fmt.Println(createDate)

	// Again this will helps to fomate the date
	fmt.Println(createDate.Format("01-02-2006"))

}

// Easy Hack - control + l to clean the terminal

/*
    Notes :-
     GOOS="darwin" ( default for the Macos )
	 GOOS="linux" ( default for the Linux )
     GOOS="windows" ( default for the Windows )
*/

// IF be habe to build for some specific system and we are on the different system  then we can use -   GOOS="linux" go build     to create the build file for that system .
