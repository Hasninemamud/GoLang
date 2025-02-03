# Arrays in Go (Golang) - Comprehensive Guide

An **array** in Go is a fixed-size, homogeneous data structure that stores elements of the same type. Unlike slices, arrays have a predefined length that cannot be changed after declaration.

---

## 1. Declaring and Initializing Arrays

### 1.1 Basic Declaration

```go
package main
import "fmt"

func main() {
    var numbers [5]int // Declares an array of 5 integers
    fmt.Println(numbers) // Output: [0 0 0 0 0]
}
```

### 1.2 Declaration with Initialization

```go
package main
import "fmt"

func main() {
    var arr = [3]int{10, 20, 30} // Initializes with values
    fmt.Println(arr) // Output: [10 20 30]
}
```

### 1.3 Using Ellipsis (`...`) for Inferred Length

```go
package main
import "fmt"

func main() {
    arr := [...]int{1, 2, 3, 4, 5} // Go infers the array length
    fmt.Println(arr) // Output: [1 2 3 4 5]
}
```

---

## 2. Accessing and Modifying Elements

```go
package main
import "fmt"

func main() {
    arr := [3]string{"Go", "Python", "Java"}
    fmt.Println(arr[1]) // Accessing the second element: Output: Python

    arr[1] = "JavaScript" // Modifying an element
    fmt.Println(arr) // Output: [Go JavaScript Java]
}
```

---

## 3. Iterating Over Arrays

### 3.1 Using `for` Loop

```go
package main
import "fmt"

func main() {
    arr := [4]int{10, 20, 30, 40}
    for i := 0; i < len(arr); i++ {
        fmt.Println(arr[i])
    }
}
```

### 3.2 Using `range`

```go
package main
import "fmt"

func main() {
    arr := [3]float64{1.1, 2.2, 3.3}
    for index, value := range arr {
        fmt.Printf("Index: %d, Value: %.1f\n", index, value)
    }
}
```

---

## 4. Multi-Dimensional Arrays

```go
package main
import "fmt"

func main() {
    matrix := [2][3]int{
        {1, 2, 3},
        {4, 5, 6},
    }
    fmt.Println(matrix) // Output: [[1 2 3] [4 5 6]]

    // Accessing elements
    fmt.Println(matrix[1][2]) // Output: 6
}
```

---

## 5. Key Properties of Arrays

- **Fixed Size:** Cannot be resized after declaration.
- **Value Type:** Assigning one array to another copies the entire data.
- **Zero-Value Initialization:** Elements are initialized with zero values by default.

---

## 6. Copying and Comparing Arrays

### 6.1 Copying Arrays

```go
package main
import "fmt"

func main() {
    a := [3]int{1, 2, 3}
    b := a // Copies all elements
    b[0] = 100
    fmt.Println(a) // Output: [1 2 3]
    fmt.Println(b) // Output: [100 2 3]
}
```

### 6.2 Comparing Arrays

```go
package main
import "fmt"

func main() {
    a := [3]int{1, 2, 3}
    b := [3]int{1, 2, 3}
    c := [3]int{4, 5, 6}

    fmt.Println(a == b) // Output: true
    fmt.Println(a == c) // Output: false
}
```

---

## Summary

- Arrays in Go are fixed-size and store elements of the same type.
- They can be single-dimensional or multi-dimensional.
- Arrays are value types, meaning they are copied when assigned to new variables.
- Use slices for more flexible data structures when dynamic resizing is needed.
