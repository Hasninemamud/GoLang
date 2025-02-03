# For Loop in Go (Golang) - Comprehensive Guide

The `for` loop in Go is the only looping construct, but it is highly versatile and can be used in various ways to handle different looping scenarios. This guide covers the different forms of `for` loops in Go.

---

## 1. Basic `for` Loop (Traditional Style)

The basic `for` loop consists of three components: **initialization**, **condition**, and **post statement (increment/decrement)**.

```go
package main
import "fmt"

func main() {
    for i := 0; i < 5; i++ {  // Initialization, Condition, Increment
        fmt.Println(i)        // Prints numbers from 0 to 4
    }
}
```

---

## 2. `for` as a `while` Loop

You can omit the initialization and post statements, effectively creating a loop that behaves like a `while` loop.

```go
package main
import "fmt"

func main() {
    j := 0
    for j < 3 {        // Condition only, similar to a while loop
        fmt.Println(j)
        j++            // Increment inside the loop
    }
}
```

---

## 3. Infinite `for` Loop

When all three components are omitted, the loop runs indefinitely. To exit the loop, you must use a `break` statement.

```go
package main
import "fmt"

func main() {
    count := 0
    for {               // Infinite loop
        fmt.Println(count)
        count++
        if count == 5 { // Breaks the loop when count equals 5
            break
        }
    }
}
```

---

## 4. `for` Loop with `range`

The `range` keyword allows you to iterate over slices, arrays, maps, strings, and channels efficiently.

```go
package main
import "fmt"

func main() {
    languages := []string{"Go", "Python", "Java"}

    for index, value := range languages {
        fmt.Println(index, value)
    }
}
```

---

## 5. Using `break` and `continue` Statements

- **`break`**: Terminates the loop immediately.
- **`continue`**: Skips the current iteration and moves to the next one.

```go
package main
import "fmt"

func main() {
    for i := 1; i <= 5; i++ {
        if i == 3 {
            continue // Skip when i is 3
        }
        if i == 5 {
            break    // Stop the loop when i equals 5
        }
        fmt.Println(i)
    }
}
```

---

## Summary

Go's `for` loop can handle a variety of use cases, including:
- Traditional `for` loops (with initialization, condition, and increment)
- `while`-like loops
- Infinite loops with controlled termination
- Iteration over collections using `range`
- Loop control using `break` and `continue`
