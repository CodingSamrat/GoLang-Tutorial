package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Go Tutorial - Defer!")

	fmt.Println("Start")
	defer fmt.Println("Middle") // This will be called just before main ends
	fmt.Println("End")


	// Multiple defer - `LIFO`
	defer fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("Third")
}
