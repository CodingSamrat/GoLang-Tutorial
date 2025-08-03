# Map In Go
In Go, a map is an unordered collection of key-value pairs. It is similar to dictionaries in Python or objects in JavaScript.
Keys must be of a comparable type (i.e., not slices or other maps), and values can be of any type.

<br/>

## Declaration
```go
// Create a map using make
myMap := make(map[string]int)

// Declare and initialize a map using shorthand syntax
countries := map[string]string{
    "IN": "India",
    "US": "America",
}
```

<br/>

## Basic Operations
Performing common tasks with maps in Go such as reading, adding, updating, and deleting elements.


### Access/Read an element
Retrieve the value associated with a key using `map[key]`. Returns the zero value if the key doesn't exist.
```go
countries := map[string]string{
    "IN": "India",
    "US": "America",
}

fmt.Println(countries["IN"])
```

> - Accessing a non-existent key returns the zero value of the map's value type.
> - Use the comma ok idiom to check for existence:

#### Example:
```go
value, ok := myMap["FOUR"]
if ok {
    fmt.Println("Exists:", value)
} else {
    fmt.Println("Key not found")
}
```

### Add or Update Element
Use `map[key] = value` to add a new key-value pair or update the value of an existing key.
```go
countries := map[string]string{
    "IN": "India",
    "US": "America",
}
countries["NP"] = "Nepal"
countries["US"] = "United State of America"
```



### Delete Element
Use `delete(map, key)` to remove a key and its associated value from the map.
```go
countries := map[string]string{
    "IN": "India",
    "US": "America",
}

// Delete an element/key
delete(countries, "US")
```

<br/>

#### Key Points
- Use `make(map[KeyType]ValueType)` to create an empty map.
- You can also use a composite literal to declare and initialize a map.
- Add or update elements using the syntax: `map[key] = value`
- Use for `key, value := range` map to iterate over the map.
- Maps are unordered, so the iteration order is not guaranteed.


<br/>

### Iterate Over a Map
Use a `for-range` loop to traverse through all `key-value` pairs in a map.

```go
for key, value := range countries {
    fmt.Printf("%v | %v\n", key, value)
}
```


<br/>

> ***NB:*** Read more on [Maps in the Go Official Documentation](https://go.dev/doc/effective_go#maps)