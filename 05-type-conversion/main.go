package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
		fmt.Println("Welcome to Go Tutorial - Type Conversion!")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a number")

	input, _ := reader.ReadString('\n')

	fmt.Printf("The number you entered %s", input)
	fmt.Printf("The type of the number is %T\n", input) // Always String

	// Parsing input (string) to Float 
	inputFloat, err := strconv.ParseFloat(strings.TrimSpace(input),64)
	
	if err!=nil {
		fmt.Println("ERROR:",err)
	}else {
		fmt.Printf("The float input => %f\n", inputFloat)
		fmt.Printf("The float type => %T\n", inputFloat)
	}


	// Parsing input (string/ASCII) to Integer 
	inputInt, err := strconv.Atoi(strings.TrimSpace(input))
	
	if err!=nil {
		fmt.Println("ERROR:",err)
	}else {
		fmt.Printf("The int input => %d\n", inputInt)
		fmt.Printf("The int type => %T\n", inputInt)
	}
}