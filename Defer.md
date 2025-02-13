# Defer in Go (Golang) - Comprehensive Guide

The `defer` statement in Go is used to postpone the execution of a function until the surrounding function returns. It is commonly used for resource cleanup, such as closing files or releasing locks.

---

## 1. Basic Syntax of `defer`

```go
package main
import "fmt"

func main() {
    defer fmt.Println("This will execute last.")
    fmt.Println("This will execute first.")
}
```

**Output:**
```
This will execute first.
This will execute last.
```

The deferred statement (`fmt.Println("This will execute last.")`) executes just before the function exits.

---

## 2. Multiple `defer` Statements

When multiple `defer` statements are present, they execute in **LIFO (Last In, First Out) order**.

```go
package main
import "fmt"

func main() {
    defer fmt.Println("First defer")
    defer fmt.Println("Second defer")
    defer fmt.Println("Third defer")
    fmt.Println("Main function execution")
}
```

**Output:**
```
Main function execution
Third defer
Second defer
First defer
```

---

## 3. Using `defer` for Resource Cleanup

### 3.1 Closing a File

```go
package main
import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close() // Ensures file is closed when function exits

    fmt.Println("Processing file...")
}
```

### 3.2 Unlocking a Mutex

```go
package main
import (
    "fmt"
    "sync"
)

func main() {
    var mu sync.Mutex
    mu.Lock()
    defer mu.Unlock() // Ensures mutex is released

    fmt.Println("Critical section")
}
```

---

## 4. Defer and Function Arguments

Deferred function calls evaluate their arguments **immediately**, but the function execution itself is delayed.

```go
package main
import "fmt"

func printValue(x int) {
    fmt.Println("Deferred value:", x)
}

func main() {
    x := 10
    defer printValue(x)
    x = 20
    fmt.Println("Updated value:", x)
}
```

**Output:**
```
Updated value: 20
Deferred value: 10
```

Even though `x` was updated to `20`, the deferred function captured its value at the time of the `defer` statement (which was `10`).

---

## 5. Defer with Anonymous Functions

```go
package main
import "fmt"

func main() {
    defer func() {
        fmt.Println("Deferred execution inside an anonymous function.")
    }()
    fmt.Println("Main function execution")
}
```

**Output:**
```
Main function execution
Deferred execution inside an anonymous function.
```

---

## Summary

- `defer` postpones execution of a function until the surrounding function exits.
- Deferred calls execute in **LIFO** order.
- Useful for **resource cleanup** (closing files, releasing locks, etc.).
- Arguments to deferred functions are evaluated immediately.
- Can be used with **anonymous functions** for inline execution.

Using `defer` effectively ensures clean and safe resource management in Go programs.
