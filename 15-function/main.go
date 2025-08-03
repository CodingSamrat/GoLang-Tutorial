package main

import (
	"fmt"
)

func main() {
	// Greet	
	greet("Welcome to Go Tutorial - Function!")


	// Add
	fmt.Println(add(9,3))


	// Divide
	value, ok := divide(9,3)
	fmt.Println(value, ok)
	

	//Sum
	fmt.Println(sum(9,3,56,64,23,2))
}

func greet(name string) {
    fmt.Println("Hey!", name)
}

func add(a int, b int) int {
    return a + b
}


func divide(a, b int) (int, bool) {
    if b == 0 {
        return 0, false
    }
    return a / b, true
}

func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}