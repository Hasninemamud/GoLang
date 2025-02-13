# Strings in Go (Golang) - Comprehensive Guide

In Go, a **string** is a sequence of bytes representing characters. Strings are immutable, meaning their contents cannot be changed after creation. This guide covers string manipulation, functions, and best practices in Go.

---

## 1. Declaring and Initializing Strings

### 1.1 Using String Literals

```go
package main
import "fmt"

func main() {
    var str1 string = "Hello, Go!"
    str2 := "Welcome to Golang"
    fmt.Println(str1)
    fmt.Println(str2)
}
```

**Output:**
```
Hello, Go!
Welcome to Golang
```

---

## 2. String Length

The `len()` function returns the number of bytes in a string.

```go
package main
import "fmt"

func main() {
    str := "Golang"
    fmt.Println("Length:", len(str))
}
```

**Output:**
```
Length: 6
```

Note: Since Go strings are UTF-8 encoded, `len()` may not reflect the number of actual characters.

---

## 3. Accessing Characters in a String

Strings can be accessed like arrays but return **byte values**.

```go
package main
import "fmt"

func main() {
    str := "Golang"
    fmt.Println("First character:", string(str[0]))
}
```

**Output:**
```
First character: G
```

---

## 4. String Concatenation

Use the `+` operator or `fmt.Sprintf` to concatenate strings.

```go
package main
import "fmt"

func main() {
    str1 := "Hello"
    str2 := "Go"
    result := str1 + ", " + str2 + "!"
    fmt.Println(result)
}
```

**Output:**
```
Hello, Go!
```

---

## 5. Multi-line Strings

Use backticks (`` ` ``) to create raw string literals.

```go
package main
import "fmt"

func main() {
    multiLine := `This is a
multi-line
string.`
    fmt.Println(multiLine)
}
```

**Output:**
```
This is a
multi-line
string.
```

---

## 6. String Comparison

Use `==` for equality checks and `strings.Compare()` for lexicographical comparison.

```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println("Go" == "Go")   // Output: true
    fmt.Println(strings.Compare("apple", "banana")) // Output: -1 (apple < banana)
}
```

---

## 7. String Functions in `strings` Package

### 7.1 Checking Substring Presence

```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(strings.Contains("Hello, Go!", "Go")) // Output: true
}
```

### 7.2 String Replacement

```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    str := "Hello, Go!"
    newStr := strings.Replace(str, "Go", "World", 1)
    fmt.Println(newStr) // Output: Hello, World!
}
```

### 7.3 Splitting a String

```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    parts := strings.Split("Go is fun", " ")
    fmt.Println(parts) // Output: [Go is fun]
}
```

### 7.4 Trimming Spaces

```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    str := "  Hello, Go!  "
    trimmed := strings.TrimSpace(str)
    fmt.Println(trimmed) // Output: Hello, Go!
}
```

---

## Summary

- Strings are immutable sequences of bytes.
- `len()` returns the byte length, not the character count.
- Use `+` or `fmt.Sprintf()` for concatenation.
- `strings` package provides utilities for searching, replacing, splitting, and trimming.
- Use raw string literals (`` ` ``) for multi-line strings.

Understanding strings in Go helps in efficient text processing and manipulation.
