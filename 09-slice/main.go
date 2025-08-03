package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Go Tutorial - Slice!")

	// Declare a slice without defining size
	var users []string

	// Append element to the slice
	users = append(users, "Sam")


	// Print the full slice
	fmt.Println("The user list =>", users)

	// Print the length of the slice
	fmt.Println("The length of user list =>", len(users))



	// Declare and initialize a slice in a single line
	var sports = []string{"Football", "Cricket",}
	sports = append(sports, "PUBG")

	// Print the initialized slice
	fmt.Println("The sports list =>", sports)

	// Print the length of the sports slice
	fmt.Println("The length of sports list =>", len(sports))


	// Sorting...
	// Declare and initialize a slice of names
	names := []string{"Zara", "Mohan", "Ankit", "Riya"}

	// Print before sorting
	fmt.Println("Before sorting:", names)

	// Sort the slice in increasing (alphabetical) order
	sort.Strings(names)

	// Print after sorting
	fmt.Println("After sorting:", names)
}