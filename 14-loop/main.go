package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Go Tutorial - Loop!")

	nums := []int{10, 20, 30}

	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}


	// Break
	for i := 1; i <= 5; i++ {
	if i == 3 {
		break // Exit loop when i is 3
	}
		fmt.Println("i =", i)
	}


	// Continue
	for i := 1; i <= 5; i++ {
	if i == 3 {
		continue // Skip when i is 3
	}
		fmt.Println("i =", i)
	}
	

	// Goto
	i := 1
	for {
		if i == 3 {
			goto end // Jump to label 'end'
		}
		fmt.Println("i =", i)
		i++
	}

	end:
	fmt.Println("Loop exited using goto")
}

