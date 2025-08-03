# Struct In Go

A struct (short for "structure") is a **user-defined type** in Go that groups together related fields to represent a real-world object. It’s like a container that holds multiple values of different types under one unit.

Structs are commonly used to model complex data like ***users, products, books,*** etc.

<br/>

## Create a Struct Type
Before using a struct in Go, you must define its structure (*user-defined type*) using the type keyword followed by a name and the struct keyword. This defines a new custom data type with grouped fields.

```go
type User struct{
	Name 		string
	Email 		string
	IsActive 	bool
}
```
This defines a new User type with three fields: `Name (string)`, `Email (string)` and `IsActive (bool)`. Think of it like creating a template or blueprint for objects.


<br/>

## Declaration
Go supports multiple ways to declare and initialize a struct. Here are the three most common:

### 1. Zero Value Declaration `var`
Declares a struct variable with zero values for all fields.Fields can be assigned individually later. Useful when you want to create an empty struct and fill it step by step.

```go
var sam User

sam.Name = "Sam"
sam.Email = "sam@codingsamrat.com"
sam.IsActive = true
```


### 2. Positional Initialization
Initializes a struct using ordered values for each field. All fields must be provided in the correct order. Less readable and fragile if the struct definition changes.

```go
jack := User{"Jack", "jack@codingsamrat.com", false}
```


### 3. Named Field Initialization
Initializes a struct using field names, regardless of order. Most readable and safe for larger structs. Recommended for better clarity and maintenance.

```go
rana := User{
    Name:     "Rana",
    Email:    "rana@codingsamrat.com",
    IsActive: true,
}
```

<br/>

> #### Notes:
> - Struct fields are accessed using the dot `(.)` operator.
> - Struct values can be initialized using:
>   - Default empty declaration `var u User`
>   - Positional literal `User{"A", "B", true, 20}`
>   - Named literal (recommended): `User{Name: "A", ...}`
> - Structs are value types, meaning they are copied when assigned or passed to functions.

<br/>

## Access individual fields
Once you have a struct variable, you can access its individual fields using the dot (`.`) notation. This allows you to retrieve specific data stored within the struct. For example, accessing the `Name` or `Age` of a `User` struct is done using `user.Name` or `user.Age`.

```go
fmt.Println("Name:", sam.Name)
fmt.Println("Is Active?", sam.IsActive)
```


<br/>

## Update value
Struct fields can be updated similarly using the dot (`.`) notation. Just assign a new value to the field using `=`. This is useful when you want to change or modify the contents of a struct after it has been initialized.

```go
sam.Name = "Samrat"
```


> ***N.B.:*** This is just a basic overview of struct operations in Go. Structs are powerful and can be used with `pointers`, `methods`, `embedding`, `tags`, and more.<br/>
> 👉 To explore advanced struct features, read the official [Go Documentation](https://golang.org/ref/spec#Struct_types)