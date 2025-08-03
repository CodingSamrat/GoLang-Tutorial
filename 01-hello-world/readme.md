# Introduction

Go (or Golang) is a modern, open-source programming language developed at Google. It is designed for simplicity, high performance, and efficient concurrency-making it a great choice for backend systems, APIs, CLI tools, and more.


## Why Learn Go?
Here’s why Go is gaining so much popularity:

- **Simple & Clean Syntax:** Easy to read, write, and maintain
- **Compiled Language:** Converts to fast native binaries (no VM)
- **Great Standard Library:** Built-in support for ***web servers, file I/O, concurrency***, etc.
- **Concurrency Built-in:** Use `goroutines` and `channels` to handle thousands of tasks efficiently
- **Cross-Platform:** Compile for ***Windows, Linux, macOS*** from a single codebase


## First Program - Hello World!
Let’s write and run your first Go program:

```go
// `/main.go`

package main

import "fmt"

func main() {
    fmt.Println("Hey, It's Coding Samrat!")
}

```

## How to Run
- Install Go (if not already): Install the go sdk from their [official website](https://go.dev/doc/install)

- Save the code as main.go

- Run the program:
```bash
go run main.go
```


## Code Breakdown
| Line                 | What it does                               |
| -------------------- | ------------------------------------------ |
| `package main`       | Declares the package name (entry point)    |
| `import "fmt"`       | Imports the `fmt` package (for formatting) |
| `func main()`        | The main function — program starts here    |
| `fmt.Println("...")` | Prints text followed by a newline          |


<br/>

> **⚠️ Note**
> - Every Go application starts with a main package and a main() function.
> - Go is strictly typed and has only one way to format things (thanks to gofmt).
