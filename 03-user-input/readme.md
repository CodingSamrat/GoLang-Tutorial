# User Input

In Go, you can take user input from the terminal using the `bufio` and `os` packages. This method is commonly used in CLI applications.

## Using `fmt.Scanln` (Standard Input)

The `fmt.Scanln()` function in Go reads input from the standard input (**keyboard**) and stores it into a variable. It stops reading at the first whitespace or newline.

```go
var input string
fmt.Print("Enter a word: ")

// Read a single word (stops at space or newline)
fmt.Scanln(&input)

fmt.Printf("You entered: %s\n", input)

// Output:
// Enter a word: Hello
// You entered: Hello
```


***Notes:***
- `fmt.Scanln(&input)` scans user input and stores it in the `input` variable.
- You must pass the address of the variable using `&` (Go assigns input to that address).
- `Scanln` stops reading at the first whitespace or newline.
- It's useful for simple inputs like numbers, single words, or multiple variables.


> **_N.B._:**  `fmt.Scanln()` doesn't handle full lines with spaces well (like full names or sentences). For reading full lines, prefer bufio.NewReader().

#### Example with Multiple Inputs:

```go
var name string
var age int

fmt.Print("Enter your name and age: ")
fmt.Scanln(&name, &age)

fmt.Printf("Name: %s, Age: %d\n", name, age)
```

---

<br/>

## Using `bufio`
The `bufio.NewReader()` reads input from standard input (keyboard), and `.ReadString('\n')` reads until the user presses **Enter**.

```go
// Create a reader to read from standard input
reader := bufio.NewReader(os.Stdin)

fmt.Print("Enter a number: ")

// Read input until newline character
input, _ := reader.ReadString('\n')

// Output the user input
fmt.Printf("The number you entered is: %s", input)

// Output:-
// Enter a number: 42
// The number you entered is: 42
```

> **_N.B._:** The line `input, _ := reader.ReadString('\n')` uses the comma ok idiom to ignore the second returned value, which is of type error. <br/>
> In real-world applications, it's best to handle the error to ensure reliable input reading.