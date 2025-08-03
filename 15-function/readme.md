# Function
Functions are reusable blocks of code that perform a specific task. In Go, every executable program must contain a `main` function, but you can define your own functions to organize your logic better, make your code reusable, and improve readability.

## Declaration
A basic function in Go is declared using the `func` keyword, followed by the function name, a list of parameters (if any), the return type (if any), and a body enclosed in `{}`.

**Go supports functions with:**

- No parameters, no return value
- Parameters and return values
- Multiple return values
- Named return values
- Variadic functions (accept variable number of arguments)

<br/>

## Function Parameters
Functions can accept any number of parameters. The type is defined after the parameter name.


```go
func greet(name string) {
    fmt.Println("Hello", name)
}
```

<br/>

## Return Values
Functions can return values using the return keyword. Go also allows multiple return values.

```go
func add(a int, b int) int {
    return a + b
}


func divide(a, b int) (int, bool) {
    if b == 0 {
        return 0, false
    }
    return a / b, true
}

```

<br/>

## Variadic Functions
You can define functions that take a variable number of arguments using `...`.

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

<br/>


> ***N.B.:*** This is just a basic introduction to functions in Go.
Read more:[ Go Functions - Official Docs](https://go.dev/doc/effective_go#functions)