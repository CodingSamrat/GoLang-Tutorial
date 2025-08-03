# Defer

In Go, the `defer` keyword is used to delay the execution of a function until the surrounding function returns.

## What does it do?
When you defer a function call, it’s scheduled to run last, right before the current function exits - even if there's a return, panic, or error.

#### Syntax
```go
defer fmt.Println("This runs at the end!")
```

You can `defer` any function call, but it’s most commonly used for:
- Closing files
- Unlocking mutexes
- Releasing network connections
- Cleanup tasks

#### Example
```go
func main() {
	fmt.Println("Start")
	defer fmt.Println("Middle") // This will be called just before main ends
	fmt.Println("End")
}

// Output: -
Start
End
Middle
```


## Multiple Defers
If you have multiple `defer` calls, they run in `LIFO` (***Last In First Out***) order:
```go
defer fmt.Println("First")
defer fmt.Println("Second")
defer fmt.Println("Third")


// Output: -
Third
Second
First
```

## Real-life Use Case
In file handling:

```go
file, err := os.Open("data.txt")
if err != nil {
	panic(err)
}
defer file.Close()
```
This ensures the file is always closed, even if an error occurs later in the function.

