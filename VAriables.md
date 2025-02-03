# Variables in Go (Golang) - Comprehensive Guide

In Go, **variables** are used to store data that can be manipulated throughout the program. This guide covers the declaration, initialization, and different ways to work with variables in Go.

---

## 1. Declaring Variables

### 1.1 Using `var` Keyword

```go
package main
import "fmt"

func main() {
    var name string = "Golang"  // Declaring with type and value
    var age int = 10            // Declaring an integer
    fmt.Println(name, age)
}
```

### 1.2 Type Inference (Without Explicit Type)

Go can infer the type based on the assigned value:

```go
package main
import "fmt"

func main() {
    var language = "Go"  // Type inferred as string
    var version = 1.18   // Type inferred as float64
    fmt.Println(language, version)
}
```

### 1.3 Short Variable Declaration (`:=`)

This is commonly used inside functions:

```go
package main
import "fmt"

func main() {
    name := "John" // Equivalent to var name string = "John"
    age := 25      // Type inferred as int
    fmt.Println(name, age)
}
```

**Note:** `:=` cannot be used outside of functions.

---

## 2. Multiple Variable Declarations

You can declare and initialize multiple variables at once:

```go
package main
import "fmt"

func main() {
    var a, b, c int = 1, 2, 3
    x, y, z := "Go", true, 3.14
    fmt.Println(a, b, c)
    fmt.Println(x, y, z)
}
```

---

## 3. Zero Values in Go

If a variable is declared without initialization, Go assigns a **zero value** based on its type:

- **Numeric types:** `0`
- **Boolean:** `false`
- **String:** `""` (empty string)
- **Pointers, slices, maps, interfaces, functions, channels:** `nil`

```go
package main
import "fmt"

func main() {
    var i int
    var s string
    var b bool
    fmt.Println(i, s, b) // Output: 0 "" false
}
```

---

## 4. Constants (`const` Keyword)

Constants are immutable values that cannot be changed after declaration:

```go
package main
import "fmt"

func main() {
    const pi = 3.14159
    const language = "Go"
    fmt.Println(pi, language)

    // pi = 3.14 // This will cause an error because constants cannot be reassigned
}
```

---

## 5. Variable Scope

- **Local Variables:** Declared inside functions and accessible only within those functions.
- **Global Variables:** Declared outside functions and accessible throughout the package.

```go
package main
import "fmt"

var globalVar = "I am global" // Global variable

func main() {
    localVar := "I am local"  // Local variable
    fmt.Println(globalVar)
    fmt.Println(localVar)
}
```

---

## 6. Blank Identifier (`_`)

The blank identifier `_` is used to ignore values you don't need:

```go
package main
import "fmt"

func main() {
    a, _ := 100, 200 // Ignore the second value
    fmt.Println(a)   // Output: 100
}
```

---

## Summary

- Use `var` for explicit variable declarations.
- Use `:=` for concise, inferred type declarations within functions.
- Variables have default zero values if uninitialized.
- Constants are immutable and defined with the `const` keyword.
- Scope determines variable accessibility within the program.

Go encourages simplicity and clarity in variable declarations to maintain clean, readable code.
