package main

import "fmt"

func main() {

	fmt.Println(" let's learn loops")

	days := []string{"Monday", "Thuesday", "Wednsesday", "Firday", "Saturday "}

	fmt.Println(days)

	// for loops
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// Another way using for loops
	for j := range days {

		fmt.Println(days[j])
	}

	// Another way using for loops
	for _, day := range days {
		fmt.Printf("index is and value is %v\n", day)
	}

	randomValue := 1

	for randomValue < 10 {

		if randomValue == 2 {
			goto lco // if be want to switch to new label when this condition get hit
		}

		if randomValue == 5 {
			randomValue++
			continue
		}
		fmt.Println(randomValue)
		randomValue++
	}

lco:
	fmt.Println("Switeched to new label")
}
