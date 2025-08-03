package main

import "fmt"

const ProjectName = "Go Tutorial" // public: when var name starts with capital letter (uppercase)
const version = "v1.0.0" // private: when var name starts with small letter (lowercase)


func main() {
	fmt.Println(ProjectName, version)
	fmt.Printf("The type of {ProjectName, version} is: %T, %T\n",ProjectName,version)

	// String
	var name string = "Sam"
	fmt.Println(name)
	fmt.Printf("The type of {name} is: %T\n",name)

	// Integer 
	var age uint8 = 26
	fmt.Println(age)
	fmt.Printf("The type of {age} is: %T\n",age)

	// Boolean
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("The type of {isLoggedIn} is: %T\n",isLoggedIn)

	// Boolean
	var gpa float32 = 9.6
	fmt.Println(gpa)
	fmt.Printf("The type of {gpa} is: %T\n",gpa)


	// Default value 
	var number int 			// The default value of `int` is 0
	fmt.Println(number)
	fmt.Printf("The type of {number} is: %T\n",number)


	// Implicit type 
	var userName ="codingsamrat"
	fmt.Println(userName)
	fmt.Printf("The type of {userName} is: %T\n",userName)


	// No var type 
	userId := 687688977			// can't be declare in global scope
	fmt.Println(userId)
	fmt.Printf("The type of {userId} is: %T\n",userId)
}