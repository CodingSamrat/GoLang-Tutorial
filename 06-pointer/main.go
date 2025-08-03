package main

import "fmt"

func main() {
fmt.Println("Welcome to Go Tutorial - Pointer!")

// Declare a pointer to an int (it holds the memory address of an int)
var address *int
fmt.Println("Initially, the pointer has no value:", address) // nil

// Declare and initialize an integer
totalPosts := 56
fmt.Println("Total posts =>", totalPosts)

// Assign the address of totalPosts to the pointer
address = &totalPosts
fmt.Println("The address (memory location) of totalPosts =>", address)

// Dereferencing the pointer to get the value it points to
fmt.Println("The total post from pointer =>", *address)

// Demonstrating both pointer and dereferenced value together
fmt.Printf("The value stored at memory location %v is %v\n", address, *address)
}