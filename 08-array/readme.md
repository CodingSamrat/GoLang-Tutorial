# Array
An array in Go is a fixed-size collection of elements of the same type. Once declared, its size cannot be changed.

```go
var users [5]string
users[0] = "Sam"
users[1] = "Jack"
users[4] = "Rana"

fmt.Println(users)          // Output: [Sam Jack     Rana]
fmt.Println(len(users))     // Output: 5
```

You can also initialize an array directly:

```go
sports := [3]string{"Football", "Cricket", "PUBG"}
fmt.Println(sports)         // Output: [Football Cricket PUBG]
fmt.Println(len(sports))    // Output: 3

```

***N.B.:***
- Arrays in Go are zero-indexed.
- You cannot access or assign beyond the defined size (e.g., users[5] will cause a compile-time error).
- Elements are automatically initialized with zero values:
    - `""` for strings
    - `0` for numbers
    - `false` for booleans
  
  
```go
var nums [4]int
fmt.Println(nums)       // [0 0 0 0]
```


### Looping through

```go
for index, value := range sports {
	fmt.Printf("Index %d: %s\n", index, value)
}
```

<br/>

> You’ll usually work with `slices`, which are more flexible, but understanding arrays first builds a solid foundation.