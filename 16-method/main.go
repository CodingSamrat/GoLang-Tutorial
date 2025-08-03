package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Go Tutorial - Method!")

	// Create a struct instance
	sam :=User{"Sam", "sam@codingsamrat.com", false}

	// Print the struct
	fmt.Println("User Name -> :", sam.Name)
	fmt.Println("User Name -> :", sam.Email)

	// Use the method of Struct
	sam.GetStatus()

}

type User struct{
	Name 		string
	Email 		string
	IsActive 	bool
}


func (u User) GetStatus() {
	if u.IsActive{
		fmt.Println("User is Active")
	}else{
		fmt.Println("User is Inactive")
	}
}
