# Variables & Datatype

## String

A **string** in Go is a sequence of bytes (characters) enclosed in double quotes (`"`). It is a **read-only slice of bytes**, meaning once a string is created, it cannot be modified (strings are immutable).

You use the `string` type to declare string variables.

``` go
var name string = "Sam"

fmt.Println(name)
fmt.Printf("The type of {name} is: %T\n",name)

// Sam
// The type of {name} is: string
```

***

<br/>

## Numeric Types

### Integer
Go provides both **signed** and **unsigned** integer types.

- **Signed integers** can represent both positive and negative numbers (e.g., `-10`, `0`, `42`).
- **Unsigned integers** only represent non-negative numbers (`0` and above), offering a larger positive range for the same bit size.

Below is a table of common **unsigned integer** types in Go:

**Unsigned Integers:**
| Type    | Size    | Range                            |
|---------|---------|----------------------------------|
| `uint8`   | 8-bit   | 0 to 255                         |
| `uint16`  | 16-bit  | 0 to 65,535                      |
| `uint32`  | 32-bit  | 0 to 4,294,967,295               |
| `uint64`  | 64-bit  | 0 to 18,446,744,073,709,551,615  | 


``` go
var age uint8 = 23

fmt.Println(age)
fmt.Printf("The type of {age} is: %T\n",age)

// 23
// The type of {age} is: uint8
```
<br/>

**Signed Integers:**


| Type    | Size    | Range                          |
|---------|---------|---------------------------------|
| `int8`  | 8-bit   | -128 to 127                    |
| `int16` | 16-bit  | -32,768 to 32,767              |
| `int32` | 32-bit  | -2,147,483,648 to 2,147,483,647|
| `int64` | 64-bit  | -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807 |



<br/>

### `int` Alias

Go also provides a general-purpose `int` type, which is an alias for either `int32` or `int64` depending on the system architecture:
- On **32-bit systems**, `int` is equal to `int32`
- On **64-bit systems**, `int` is equal to `int64`

Use `int` when you don’t need a specific bit size — it’s idiomatic and simplifies your code.


```go
var number int      // Default value is 0
fmt.Println(number)

// 0
```

<br/>

### Float

A **float** (floating-point number) in Go is used to represent real numbers (i.e., numbers with a decimal point). Go provides two main floating-point types:

- `float32`: 32-bit precision
- `float64`: 64-bit precision (default and more precise)

Use `float64` unless you have a specific reason to use less memory.

``` go
var price float32 = 99.99
var rating float64 = 4.789

fmt.Println("Price:", price)
fmt.Println("Rating:", rating)
fmt.Printf("Type of price: %T\n", price)
fmt.Printf("Type of rating: %T\n", rating)

// Price: 99.99
// Rating: 4.789
// Type of price: float32
// Type of rating: float64
```

***Notes***:

- Floats can hold decimal values, unlike integers.
- Default type when assigning a decimal number like 3.14 is float64.
- Use `%f` in `Printf` to format float output (e.g., to limit decimal places).

Example with Formatting:
``` go
fmt.Printf("Rating: %.2f\n", rating)  // Output: Rating: 4.79
```

***

<br/>

## Boolean

A **boolean** type in Go, written as `bool`, represents a value that can be either `true` or `false`.

Booleans are commonly used in:
- Conditional statements (`if`, `for`)
- Logical operations (`&&`, `||`, `!`)


```go
var isLoggedIn bool = true
var isAdmin = false     // type inferred as bool

fmt.Println("Logged In:", isLoggedIn)
fmt.Println("Is Admin:", isAdmin)
fmt.Printf("Type of isLoggedIn: %T\n", isLoggedIn)

// Logged In: true
// Is Admin: false
// Type of isLoggedIn: bool
```

### Default Value

If a `bool` is declared without a value, it defaults to `false`.

``` go
var active bool
fmt.Println(active)     // Output: false
```



***

<br/>

## Short Variable Declaration (`:=`)

Go allows you to declare and initialize variables quickly using the `:=` syntax. This is called **short variable declaration**.

```go
userId := 687688977
fmt.Println(userId)
fmt.Printf("The type of {userId} is: %T\n", userId)

// 687688977
// The type of {userId} is: int
```

**Notes:**

- You must use `:=` inside a function body (usually `main()` or any other function).
- Go infers the type automatically from the assigned value.
- This is the most commonly used way to declare variables in Go, especially for temporary or `local` use.


<br/>

**Important:**

``` go
userId := 12345  // ✅ valid inside function
userId := 12345  // ❌ invalid outside function scope
```

Use `var` or `const` for global declarations instead:
``` go
var userId = 12345  // ✅ valid globally
```