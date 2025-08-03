package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Go Tutorial - Switch Case!")

	day := "Sunday"

	switch day {
	case "Monday":
		fmt.Println("Let's go to work")
	case "Saturday", "Sunday":
		fmt.Println("Chill babe! It's weekend...")
	default:
		fmt.Println("More work...")
	}
	
}

