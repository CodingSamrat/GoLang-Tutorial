# Method
In Go, methods are functions that are associated with a type, typically a `struct`. They allow you to define behavior specific to that type, making your code more modular and object-oriented in style.

## Declaring a Method
To declare a method, use a receiver; a special parameter that appears before the function name. It ties the function to a specific type:

```go
func (u User) MethodName() {
	// method logic
}

```
Here, `u User` is the receiver. It means `MethodName()` is a method on the User type.


#### Example:

```go
func main() {
	// Create a struct instance
	sam :=User{"Sam", "sam@codingsamrat.com", false}

	sam.GetStatus()
}

// The Struct
type User struct{
	Name 		string
	Email 		string
	IsActive 	bool
}

// The method of struct - User
func (u User) GetStatus() {
	if u.IsActive{
		fmt.Println("User is Active")
	}else{
		fmt.Println("User is Inactive")
	}
}
```

<br/>

### Value vs Pointer Receiver
- Value receiver `(u User)` receives a copy of the struct.
- Pointer receiver` (u *User)` allows modification of the original struct.

**Use pointer receivers when:**
- You want to modify the original data.
- The struct is large and copying it is expensive.

> ***Note:*** In Go, methods give you a clean and intuitive way to add behavior to your types without using traditional OOP-style classes.

