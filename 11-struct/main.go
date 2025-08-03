package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Go Tutorial - Struct!")

	// Create a struct instance
	var sam User
	sam.Name = "Sam"
	sam.Email = "sam@codingsamrat.com"
	sam.IsActive = true

	// Print the struct
	fmt.Println("User details:", sam)

	// Update value
	sam.Name = "Samrat"
	fmt.Println("User details after update:", sam)

	// Access individual fields
	fmt.Println("Name:", sam.Name)
	fmt.Println("Is Active?", sam.IsActive)
	

	// Another way: struct literal
	jack := User{"Jack", "jack@codingsamrat.com", false}
	fmt.Println("User:", jack)


	// Named fields (recommended for readability)
	rana := User{
		Name:     "Rana",
		Email:    "rana@codingsamrat.com",
		IsActive: true,
	}
	fmt.Println("User:", rana)
}

type User struct{
	Name 		string
	Email 		string
	IsActive 	bool
}