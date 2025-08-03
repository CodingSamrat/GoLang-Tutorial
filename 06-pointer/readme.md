# Pointers in Go

In Go, pointers allow you to store and work with the memory address of a variable, instead of the actual value. This is useful for passing large data structures efficiently or modifying a value inside a function.

## Declaration
To declare a pointer:


```go
var ptr *int
```
- `int` means it's a pointer to an `int`.
- A pointer stores the **address** of a value, not the value itself.


## Getting the Address (&)
The & operator is used to get the address of a variable.

```go
x := 42
ptr := &x
```
Now `ptr` holds the memory address of `x`.


## Dereferencing (*)
The `*` operator is used to access the value stored at a memory address (i.e., the value the pointer points to).

```go
fmt.Println(*ptr) // prints 42
```

### Example

```go
fmt.Println("Welcome to Go Tutorial - Pointer!")

// Declare a pointer to an int (it holds the memory address of an int)
// Syntax - `var ptr * type`
var address *int
fmt.Println("Initially, the pointer has no value:", address) // nil

// Declare and initialize an integer
totalPosts := 56
fmt.Println("Total posts =>", totalPosts)

// Assign the address of totalPosts to the pointer
address = &totalPosts
fmt.Println("The address (memory location) of totalPosts =>", address)

// Dereferencing the pointer to get the value it points to
fmt.Println("The total post from pointer =>", *address)

// Demonstrating both pointer and dereferenced value together
fmt.Printf("The value stored at memory location %v is %v\n", address, *address)
```


#### Note
- Go does not support pointer arithmetic (like **C/C++**).
- You can use pointers to modify values inside a function.
- Default value of a pointer is `nil` if not initialized.



## Passing Pointers to Functions
In Go, **functions receive copies** of their arguments. If you want a function to modify the original value, you need to pass a **pointer** to that value.

### Why use pointers?
By passing pointer:
- Avoids copying large structures
- Enables the function to change the original value

```go
func updateLikes(likes *int) {
    *likes += 1
}

func main() {
    totalLikes := 100
    fmt.Println("Before:", totalLikes)

    // No need any return value or reassigning to another variable
    // It update the value of memory location of `totalLikes`
    updateLikes(&totalLikes)

    fmt.Println("After:", totalLikes)
}


// Output
// Before: 100
// After: 101
```


#### How it works:
- `updateLikes(likes *int)` takes a pointer to an `int`.
- `&totalLikes` passes the address of `totalLikes`.
- Inside the function, `*likes` dereferences the pointer to access and modify the original value.

#### ⚠️ Notes:
- Changes made via the pointer are reflected outside the function.
- Useful when dealing with large data structures or wanting to avoid value copies.