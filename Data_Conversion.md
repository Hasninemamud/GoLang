# Data Conversion in Go (Golang) - Comprehensive Guide

Data conversion (also known as type casting) in Go is the process of converting a value from one data type to another. Go does not support implicit type conversions, meaning all conversions must be done explicitly.

---

## 1. Basic Type Conversion

Go provides a simple syntax for converting one type to another:

```go
variable := type(value)
```

### Example:

```go
package main
import "fmt"

func main() {
    var i int = 42
    var f float64 = float64(i) // Converting int to float64
    var u uint = uint(f)       // Converting float64 to uint
    
    fmt.Println(i, f, u)
}
```

**Output:**
```
42 42.000000 42
```

---

## 2. String to Integer Conversion

Use `strconv.Atoi()` to convert a string to an integer.

```go
package main
import (
    "fmt"
    "strconv"
)

func main() {
    str := "123"
    num, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println("Conversion error:", err)
    } else {
        fmt.Println("Converted Integer:", num)
    }
}
```

**Output:**
```
Converted Integer: 123
```

---

## 3. Integer to String Conversion

Use `strconv.Itoa()` to convert an integer to a string.

```go
package main
import (
    "fmt"
    "strconv"
)

func main() {
    num := 456
    str := strconv.Itoa(num)
    fmt.Println("Converted String:", str)
}
```

**Output:**
```
Converted String: 456
```

---

## 4. String to Float Conversion

Use `strconv.ParseFloat()` to convert a string to a float.

```go
package main
import (
    "fmt"
    "strconv"
)

func main() {
    str := "3.14"
    f, err := strconv.ParseFloat(str, 64)
    if err != nil {
        fmt.Println("Conversion error:", err)
    } else {
        fmt.Println("Converted Float:", f)
    }
}
```

**Output:**
```
Converted Float: 3.14
```

---

## 5. Float to String Conversion

Use `strconv.FormatFloat()` to convert a float to a string.

```go
package main
import (
    "fmt"
    "strconv"
)

func main() {
    f := 9.876
    str := strconv.FormatFloat(f, 'f', 2, 64) // 'f' means decimal notation, 2 decimal places
    fmt.Println("Converted String:", str)
}
```

**Output:**
```
Converted String: 9.88
```

---

## 6. Boolean to String Conversion

Use `strconv.FormatBool()` to convert a boolean to a string.

```go
package main
import (
    "fmt"
    "strconv"
)

func main() {
    b := true
    str := strconv.FormatBool(b)
    fmt.Println("Converted String:", str)
}
```

**Output:**
```
Converted String: true
```

---

## 7. String to Boolean Conversion

Use `strconv.ParseBool()` to convert a string to a boolean.

```go
package main
import (
    "fmt"
    "strconv"
)

func main() {
    str := "true"
    b, err := strconv.ParseBool(str)
    if err != nil {
        fmt.Println("Conversion error:", err)
    } else {
        fmt.Println("Converted Boolean:", b)
    }
}
```

**Output:**
```
Converted Boolean: true
```

---

## 8. Summary

- **Explicit type conversion** is required in Go (`type(value)`).
- Use `strconv` package for converting between **strings and other types**.
- Common conversion functions:
  - `strconv.Atoi()` - String to Integer
  - `strconv.Itoa()` - Integer to String
  - `strconv.ParseFloat()` - String to Float
  - `strconv.FormatFloat()` - Float to String
  - `strconv.FormatBool()` - Boolean to String
  - `strconv.ParseBool()` - String to Boolean

Understanding data conversion in Go ensures efficient type handling and avoids common runtime errors.
