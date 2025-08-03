## GoLang Tutorial
Welcome to the ***Go Language Tutorial*** – a beginner-friendly guide to learning Go (also known as Golang), a modern programming language designed for simplicity, performance, and reliability.

This tutorial covers the essential concepts of Go with simple, well-documented examples. Each topic includes explanations and code snippets to help you understand how things work in real-world scenarios.


## Topics Covered
- Getting Started (Hello World)
- Variables and Constants
- Data Types
- Type Casting
- Pointers
- Arrays and Slices
- Maps
- Structs
- Conditional Statements (`if`, `else`, `switch`)
- Loops (`for`, `range`)
- Control Statements (`break`, `continue`, `goto`)
- Functions
- Methods
- Defer
- File Handling


## Getting Started
To run the examples, make sure you have Go installed:
```bash
go version
```

Clone this repository, visit the specific dir (topic) and run any file using:
```bash
go run main.go
```


#### Here's a sample directory structure:
```plaintext
/GoLang-Tutorial
│
├── README.md
├── readme.md
│
├── 01-hello-world            # Topic folder: Hello World
│   ├── README.md             # Topic-specific doc and explanation
│   ├── go.mod                # Module-specific file (optional, if standalone)
│   └── main.go               # Main Go file with example code
└── 02- ...
```


<br/>

## Build and Run Go Programs
In Go, you can build and run your programs using the built-in go command-line tool. It compiles source code into an executable binary.
```bash
go build

# Build with executable name
go build -o app_name
```
