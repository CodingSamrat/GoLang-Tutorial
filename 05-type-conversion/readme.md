# Type Conversion

In Go, type  conversion allows you to convert a variable from one type to another explicitly. Unlike some languages, Go does not support implicit type conversions - you must do it manually.

## Why Type Conversion Needed?
You might need to convert types when:

- Doing arithmetic between different numeric types.
- Converting a number to a string or a numeric string to number. 
- Parsing user input or api params (which is read as a string) into a number.


## Basic Syntax
The basic syntax of type conversion is `TargetType(value)`.
```go
var a int = 42
var b float64 = float64(a)
```

---

## Numeric Type Conversion

```go
var a int = 10
var b float64 = float64(a)    // Cast int to float64
var c int = int(b)            // Cast float64 to int

fmt.Println("a:", a)
fmt.Println("b:", b)
fmt.Println("c:", c)
```

---

## String to Number (Parsing)
When reading input or converting strings, use the `strconv` package:

```go
str := "123"                    // numeric string
num, err := strconv.Atoi(str)   // Convert string to int

if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Parsed number:", num)
}
```

---

## Number to String
```go
num := 42
str := strconv.Itoa(num)      // Convert int to string
fmt.Println("Number as string:", str)
```

> ***N.B.:*** `strconv.Atoi(...)` stands for ASCII to int and `strconv.Itoa(...)` for vice versa




## Some Useful `strconv` Functions

| Function                    | Purpose                         |
| --------------------------- | --------------------------------|
| `strconv.Atoi(s)`           | Convert string (ASCII) to int   |
| `strconv.ParseFloat(s, 64)` | Convert string to float64       |
| `strconv.Itoa(n)`           | Convert int to string (ASCII)   |
| `fmt.Sprintf("%T", val)`    | Get the type of a variable      |


<br/>

#### Notes:
- You must convert explicitly, **Go** will not auto-convert like **JavaScript** or **Python**.
- Type casting only works between compatible types.
- For input, use functions like `strconv.Atoi()`, `strconv.ParseFloat()`, etc.
