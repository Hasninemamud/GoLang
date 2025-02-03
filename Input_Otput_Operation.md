# Input and Output Operations in Go (Golang) - Comprehensive Guide

This guide covers essential functions for output (`Println`, `Printf`, `Print`) and taking user input in Go. These functions are part of the `fmt` package, which provides formatted I/O with functions analogous to C's `printf` and `scanf`.

---

## 1. Printing Output in Go

### 1.1 Using `fmt.Println`

`Println` prints output with a newline at the end.

```go
package main
import "fmt"

func main() {
    fmt.Println("Hello, Go!") // Adds a new line automatically
    fmt.Println("Line 1", "Line 2") // Prints with a space between values
}
```

**Output:**
```
Hello, Go!
Line 1 Line 2
```

---

### 1.2 Using `fmt.Printf`

`Printf` allows formatted output, similar to C's `printf`.

```go
package main
import "fmt"

func main() {
    name := "John"
    age := 30
    fmt.Printf("Name: %s, Age: %d\n", name, age)
}
```

- `%s` - String
- `%d` - Integer
- `%f` - Float
- `\n` - Newline character

**Output:**
```
Name: John, Age: 30
```

You can also control floating-point precision:

```go
pi := 3.14159
fmt.Printf("Pi: %.2f\n", pi) // Output: Pi: 3.14
```

---

### 1.3 Using `fmt.Print`

`Print` prints without adding a newline or space automatically.

```go
package main
import "fmt"

func main() {
    fmt.Print("Hello, ")
    fmt.Print("World!")
}
```

**Output:**
```
Hello, World!
```

---

## 2. Taking User Input

Go provides the `fmt.Scan`, `fmt.Scanf`, and `fmt.Scanln` functions for reading input from the user.

### 2.1 Using `fmt.Scan`

Reads space-separated input values.

```go
package main
import "fmt"

func main() {
    var name string
    var age int
    fmt.Print("Enter your name and age: ")
    fmt.Scan(&name, &age) // Use & to pass variable addresses
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
}
```

**Example Input:**
```
John 25
```

**Output:**
```
Name: John
Age: 25
```

---

### 2.2 Using `fmt.Scanln`

Reads input until a newline is encountered (better for single-line inputs).

```go
package main
import "fmt"

func main() {
    var name string
    fmt.Print("Enter your name: ")
    fmt.Scanln(&name)
    fmt.Println("Hello,", name)
}
```

**Example Input:**
```
Alice
```

**Output:**
```
Hello, Alice
```

---

### 2.3 Using `fmt.Scanf`

Reads formatted input based on specified format specifiers.

```go
package main
import "fmt"

func main() {
    var name string
    var age int
    fmt.Print("Enter your name and age (e.g., John 30): ")
    fmt.Scanf("%s %d", &name, &age)
    fmt.Printf("Name: %s, Age: %d\n", name, age)
}
```

**Example Input:**
```
John 30
```

**Output:**
```
Name: John, Age: 30
```

---

## Summary

- **Output Functions:**
  - `fmt.Println` - Prints with a newline.
  - `fmt.Printf` - Prints formatted text.
  - `fmt.Print` - Prints without a newline.

- **Input Functions:**
  - `fmt.Scan` - Reads space-separated input.
  - `fmt.Scanln` - Reads until a newline.
  - `fmt.Scanf` - Reads formatted input based on format specifiers.

Use these functions effectively for handling standard I/O operations in Go.
