<aside>
💡 In Go, a variable is a storage location that holds a value of a specific data type. Variables are fundamental to programming and allow you to store and manipulate data within your programs. Here's how you declare and use variables in Go:

1. **Declaration:**
    
    To declare a variable in Go, you use the `var` keyword, followed by the variable name, the data type, and optionally, an initial value. For example:
    
    ```go
    var age int
    var name string
    var isStudent bool
    
    ```
    
    You can also declare multiple variables of the same data type in a single line:
    
    ```go
    var x, y int
    
    ```
    
    If you want to declare and initialize a variable in one step, you can do so like this:
    
    ```go
    var score float64 = 98.5
    var message string = "Hello, world!"
    
    ```
    
    In many cases, Go can infer the data type from the assigned value, so you can omit the data type:
    
    ```go
    var age = 25 // Go infers that age is of type int
    
    ```
    
    You can also use the short variable declaration `:=` within a function to declare and initialize variables:
    
    ```go
    age := 30 // shorthand for "var age int = 30"
    
    ```
    
2. **Assignment:**
    
    You can assign values to variables using the `=` operator. For example:
    
    ```go
    age = 35
    name = "Alice"
    isStudent = true
    
    ```
    
3. **Variable Naming Rules:**
    - Variable names must begin with a letter (a-z, A-Z) or an underscore (_) and can be followed by letters, digits (0-9), or underscores.
    - Go is case-sensitive, so `myVar` and `myvar` are treated as different variables.
    - Avoid using reserved keywords like `int`, `string`, `var`, etc., as variable names.
4. **Scope:**
    
    The scope of a variable determines where in your code it can be accessed. In Go, variables can have package-level scope or function-level scope, depending on where they are declared.
    
    - Package-level variables are declared at the package level and can be accessed throughout the package.
    - Function-level variables are declared within a function and can only be accessed within that function.

Here's an example of declaring and using variables in Go:

```go
package main

import "fmt"

func main() {
    var age int // Declaration
    age = 25     // Assignment

    var name string = "John" // Declaration and initialization

    isStudent := true // Short variable declaration

    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Is Student:", isStudent)
}

```

This code declares and uses three variables (`age`, `name`, and `isStudent`) of different data types and prints their values to the console.

</aside>