package main

import "fmt"

func main() {
	fmt.Println("Welcome to Go Tutorial - Array!")

	// Declare an array of 5 strings
	var users [5]string

	// Assign values to specific indexes
	users[0] = "Sam"
	users[1] = "Jack"
	users[4] = "Rana"

	// users[5] = "Error!" // ❌ Compile-time error: index out of range

	// Print the full array
	fmt.Println("The user list =>", users)

	// Print the length of the array
	fmt.Println("The length of user list =>", len(users))

	// Declare and initialize an array in a single line
	var sports = [3]string{"Football", "Cricket", "PUBG"}

	// Print the initialized array
	fmt.Println("The sports list =>", sports)

	// Print the length of the sports array
	fmt.Println("The length of sports list =>", len(sports))
}