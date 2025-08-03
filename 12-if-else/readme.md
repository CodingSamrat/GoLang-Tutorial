# If-Else In Go

In Go, conditional logic is implemented using `if`, `else if` and `else` statements. These are used to execute different blocks of code based on boolean expressions.

```go
age := 17

if age >= 18 {
	fmt.Println("You are allowed to vote")
}else{
	fmt.Println("You are not allowed to vote")
}
```
This structure allows programs to branch logic based on conditions, which is a fundamental concept in programming.

> ***N.B.:*** In Go, parentheses around the condition are optional (unlike some other languages), but curly braces `{}` are required.



## Conditional Operators in Go
Here is the list of condition operators used in Go. Assuming `a = 10`, `b = 5`.

| Operator | Description              | Example     | Result  |
| -------- | ------------------------ | ----------- | ------- |
| `==`     | Equal to                 | `a == b`    | `false` |
| `!=`     | Not equal to             | `a != b`    | `true`  |
| `>`      | Greater than             | `a > b`     | `true`  |
| `<`      | Less than                | `a < b`     | `false` |
| `>=`     | Greater than or equal to | `a >= 10`   | `true`  |
| `<=`     | Less than or equal to    | `b <= 5`    | `true`  |


These operators are used to compare values and return a boolean (`true` or `false`), which can control the program flow using `if`, `for`, etc.

> **Tip:** Combine them with logical operators like `&&` (AND), `||` (OR), and `!` (NOT) for more complex conditions.

#### Another Example
```go
if 3%2 ==0{
    fmt.Println("Even Number")
}else{
    fmt.Println("Odd Number")
}
```

<br/>

### Special if Statement with Short Variable Declaration
Go allows you to declare a variable inside the if statement using a short declaration (`:=`). The variable is scoped only within the `if-else` block.

```go
if num := 2; num>10{
    fmt.Println("Number is grater than 10")
}else if num == 10{
    fmt.Println("Number is equal to 10")
}else{
    fmt.Println("Number is less than 10")
}
```

#### Explanation:
- `num := 2` is declared within the `if` block and is only accessible inside the `if`, `else if`, and `else`.
- This pattern is useful when you only need a variable temporarily for conditional checks.

> ***N.B.:*** The variable declared in this form is not accessible outside the if block. It's ideal for short-lived checks.