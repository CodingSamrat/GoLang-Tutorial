# Slice In Go-Lang

A slice in Go is a dynamic, flexible view of an underlying array. It’s the most commonly used data structure for handling lists.

<br/>

## Declaration
```go
// Empty slice
var names []string

// Slice declaration with initial value 
var names []string{"sam", "Jordy"}
```


<br/>

## Append value
You can append values using the `append()` function:

```go
names = append(names, "Sam")
```


Multiple values can be also added at once.

```go
colors := []string{"Red"}
colors = append(colors, "Green", "Blue")
```



#### Example:

```go
sports := []string{"Football", "Cricket"}
sports = append(sports, "PUBG")

fmt.Println(sports)         // Output: [Football Cricket PUBG]
fmt.Println(len(sports))    // Output: 3
```


<br/>

### Appending One Slice to Another in Go
Learn how to merge two slices using the `append()` function and the variadic `...` operator in Go.

```go
fruits := []string{"Apple", "Banana"}
moreFruits := []string{"Mango", "Orange"}

// Append moreFruits slice to the fruits slice
fruits = append(fruits, moreFruits...)  // Note the '...'

fmt.Println("All Fruits:", fruits)
```

> The `...` operator is essential when appending one slice to another in Go. It's called the ***variadic expansion***.



<br/>

##  Slicing 
In Go, slicing refers to creating a new view (slice) over an existing array or slice using the `[low:high]` syntax:

```go
data := []string{"a", "b", "c", "d"}

slice1 := data[1:3]     // Includes index 1 and 2 → ["b", "c"]
slice2 := data[:2]      // From beginning to index 1 → ["a", "b"]
slice3 := data[2:]      // From index 2 to end → ["c", "d"]
```



<br/>

##  Remove Value From Slice 
In Go, slices are dynamic but don’t have built-in methods like `.remove()` (as in _Python_ or _JS_). So you need to remove elements manually using slicing or append. Here is how we can do this.

### 1. Remove by Index
```go
// Remove the element at index 1
nums := []int{10, 20, 30, 40}
indexToRemove := 1

nums = append(nums[:indexToRemove], nums[indexToRemove+1:]...)
fmt.Println(nums)       // Output: [10 30 40]
```


### 2. Remove First Element
```go
data := []string{"a", "b", "c"}
data = data[1:]         // keeps from index `1` to rest
fmt.Println(data)       // Output: [b c]
```


### 3. Remove Last Element
```go
data := []string{"a", "b", "c"}
data = data[:len(data)-1]   // keeps from index `0` to `n-1`
fmt.Println(data)           // Output: [b c]
```


### 4. Remove by Value (manual loop)
```go
names := []string{"Sam", "Jack", "Rana"}
toRemove := "Jack"
result := []string{}

for _, name := range names {
	if name != toRemove {
		result = append(result, name)
	}
}

fmt.Println(result)     // Output: [Sam Rana]
```

<br/>

> - Slice removal in Go is not built-in because Go prefers simplicity and transparency in memory usage. 
> - Always handle edge cases like invalid indexes or empty slices when removing.
> - If memory reuse is important, avoid leaving unused references in the underlying array.


<br/>

## Sorting a Slice in Go
In Go, the standard library provides the `sort` package to sort slices. You can use functions like `sort.Strings()` for strings and `sort.Ints()` for integers.

```go
names := []string{"Zara", "Mohan", "Ankit", "Riya"}

fmt.Println("Before sorting:", names)
sort.Strings(names)
fmt.Println("After sorting:", names)


// Output:
// Before sorting: [Zara Mohan Ankit Riya]
// After sorting: [Ankit Mohan Riya Zara]
```



<br/>

##  Slice vs Array

| Feature        | Array         | Slice                    |
| -------------- | ------------- | ------------------------ |
| Size           | Fixed         | Dynamic (resizable)      |
| Declaration    | `[5]string{}` | `[]string{}`             |
| Memory usage   | Pre-allocated | Based on usage           |
| Append support | ❌ Not allowed | ✅ Allowed via `append()` |


<br/>

***N.B.:***
- Slices internally point to an underlying array, so multiple slices can share the same backing array.
- Appending may create a new underlying array if capacity is exceeded.
- Use `len()` for current number of elements and `cap()` for capacity.

