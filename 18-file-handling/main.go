package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to Go Tutorial - File Handling!")

	filePath := "./database.txt"
	content := "this is the file content"

	writeFile(filePath, content) // Write content to file
	readFile(filePath)           // Read and print content from file
}



func writeFile(path string, content string) {
	file, err := os.Create(path) // Create or truncate the file
	if err != nil {
		panic(err) // Handle file creation error
	}

	length, err := io.WriteString(file, content) // Write content as string
	if err != nil {
		panic(err) // Handle write error
	}

	fmt.Println("Length -", length)

	defer file.Close() // Always close the file to release resources
}


func readFile(path string) {
	dataByte, err := os.ReadFile(path) // Read file content as byte slice
	if err != nil {
		panic(err) // Handle read error
	}

	fmt.Println("Data Byte:", dataByte)             // Print raw bytes
	fmt.Println("Data String:", string(dataByte))   // Convert bytes to string
}
