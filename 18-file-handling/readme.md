# File Handling
File handling in Go is straightforward and efficient. This example demonstrates how to create, write to, and read from a file using Go's standard os and io packages.

## Writing to a File
The writeFile function creates a file and writes a string to it:

```go
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
```

### Key Points:
- `os.Create(path)`: Creates or truncates the file at the given path.
- `io.WriteString(file, content)`: Writes a string to the file.
- `defer file.Close()`: Ensures the file is closed after the function exits.


## Reading from a File
The readFile function reads and prints the file contents:

```go
func readFile(path string) {
	dataByte, err := os.ReadFile(path) // Read file content as byte slice
	if err != nil {
		panic(err) // Handle read error
	}

	fmt.Println("Data Byte:", dataByte)             // Print raw bytes
	fmt.Println("Data String:", string(dataByte))   // Convert bytes to string
}
```

### Key Points:
- `os.ReadFile(path)`: Reads the entire file and returns content as a byte slice.
- `string(dataByte)`: Converts bytes to readable string.

<br/>
<br/>

## Main Function
```go
func main() {
	fmt.Println("Welcome to Go Tutorial - File Handling!")

	filePath := "./database.txt"
	content := "this is the file content"

	writeFile(filePath, content) // Write content to file
	readFile(filePath)           // Read and print content from file
}
```

#### Output
```txt
Welcome to Go Tutorial - File Handling!
Length - 25
Data Byte [116 104 105 115 32 105 115 32 116 104 101 32 102 105 108 101 32 99 111 110 116 101 110 116]
Data String this is the file content
```

> ***Note:*** Always use defer file.Close() after opening or creating a file to avoid memory/resource leaks.

> ***N.B.:*** This is a basic example of file I/O. To learn more about advanced file operations like appending, file permissions, or buffered I/O, check the official documentation: [Go File I/O Docs](https://pkg.go.dev/os)