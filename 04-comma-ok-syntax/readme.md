# Comma ok syntax `_, ok`:
In Go, functions often return **two values** - the result and a second value that indicates success or failure. This is commonly referred to as the **comma-ok** syntax or **value-error** handling pattern.

In the example below, the second return value is an `error`, which should **always be checked** in real-world applications.


```go
inputFloat, err := strconv.ParseFloat(strings.TrimSpace(input),32)
	
if err!=nil {
    fmt.Println("ERROR:",err)
}else {
    fmt.Printf("The float input => %f\n", inputFloat)
}
```
> **_N.B._:** In this example we used `strconv.ParseFloat` & `strings.TrimSpace`, which is type casting. Use here to simulate and catch the error. We will learn about it later.


