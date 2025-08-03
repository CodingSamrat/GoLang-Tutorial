# Loop
Go has only one looping construct: the `for` loop. It’s versatile and can be used in various forms — including ***traditional loops, while-like loops,*** and ***infinite loop***.

## Basic for Loop
This is the most common format:

```go
for i := 0; i < 5; i++ {
    fmt.Println("Count:", i)
}
```

<br/>

## `for` as a While Loop
If you omit the init and post statements, for behaves like a while loop:

```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```

<br/>

## Infinite Loop
If you omit all parts of the `for` statement, it runs forever:

```go
for {
    fmt.Println("Looping forever...")
}
```
You can break it using `break`, or skip iterations using `continue`.


<br/>

## Looping over Collections (`for-range`)
Used to iterate over *slices, arrays, maps, strings,* and *channels*:

```go
nums := []int{10, 20, 30}

for index, value := range nums {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

> ***N.B.:*** Go does not have a separate `while` or `do-while` construct. The for loop covers all looping needs.


## Control Flow Keywords in Go
Go provides a few powerful statements to control loop execution:

### Continue
The `continue` statement skips the current iteration and moves to the next one.

```go
for i := 1; i <= 5; i++ {
	if i == 3 {
		continue // Skip when i is 3
	}
	fmt.Println("i =", i)
}
```

### Break
The `break` statement immediately stops the loop.

```go
for i := 1; i <= 5; i++ {
	if i == 3 {
		break // Exit loop when i is 3
	}
	fmt.Println("i =", i)
}
```


### Goto
The `goto` statement jumps to a labeled statement. It's rarely used but helpful in special scenarios.

```go
i := 1
for {
	if i == 3 {
		goto end // Jump to label 'end'
	}
	fmt.Println("i =", i)
	i++
}

end:
fmt.Println("Loop exited using goto")
```

> ***N.B.:*** While `break` and `continue` are common in many languages, `goto` is considered a low-level control flow tool. Use it only when necessary, like in deeply nested loops or error handling patterns.