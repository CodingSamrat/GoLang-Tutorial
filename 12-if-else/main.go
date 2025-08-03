package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Go Tutorial - If-Else!")

	age := 17

	if age >= 18 {
		fmt.Println("You are allowed to vote")
	}else{
		fmt.Println("You are not allowed to vote")
	}
	

	// Another Example
	if 3%2 ==0{
		fmt.Println("Even Number")
	}else{
		fmt.Println("Odd Number")
	}



	// Special Syntax
	if num := 2; num>10{
		fmt.Println("Number is grater than 10")
	}else if num == 10{
		fmt.Println("Number is equal to 10")
	}else{
		fmt.Println("Number is less than 10")
	}
}

