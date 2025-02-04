# Maps in Go (Golang) - Comprehensive Guide

A **map** in Go is an unordered collection of key-value pairs. Maps are used when you need to associate values with unique keys, similar to dictionaries in Python or objects in JavaScript.

---

## 1. Declaring and Initializing Maps

### 1.1 Using `make()` Function

```go
package main
import "fmt"

func main() {
    m := make(map[string]int) // Creates an empty map with string keys and int values
    m["apple"] = 5
    m["banana"] = 3

    fmt.Println(m) // Output: map[apple:5 banana:3]
}
```

### 1.2 Using Map Literals

```go
package main
import "fmt"

func main() {
    m := map[string]string{
        "name":    "John",
        "country": "USA",
    }
    fmt.Println(m) // Output: map[name:John country:USA]
}
```

---

## 2. Accessing Map Elements

```go
package main
import "fmt"

func main() {
    m := map[string]int{"apple": 5, "banana": 3}
    fmt.Println(m["apple"]) // Accessing value by key: Output: 5

    // Accessing a non-existent key returns the zero value of the value type
    fmt.Println(m["grape"]) // Output: 0
}
```

---

## 3. Checking if a Key Exists

You can use the **comma ok** idiom to check if a key exists in the map:

```go
package main
import "fmt"

func main() {
    m := map[string]int{"apple": 5, "banana": 3}
    value, exists := m["banana"]
    if exists {
        fmt.Println("Value:", value)
    } else {
        fmt.Println("Key does not exist")
    }
}
```

**Output:**
```
Value: 3
```

---

## 4. Updating Map Values

```go
package main
import "fmt"

func main() {
    m := map[string]int{"apple": 5, "banana": 3}
    m["apple"] = 10 // Updating the value for an existing key
    fmt.Println(m)  // Output: map[apple:10 banana:3]
}
```

---

## 5. Deleting Map Elements

Use the `delete()` function to remove a key-value pair from the map:

```go
package main
import "fmt"

func main() {
    m := map[string]int{"apple": 5, "banana": 3}
    delete(m, "banana") // Removes the key "banana"
    fmt.Println(m)      // Output: map[apple:5]
}
```

---

## 6. Iterating Over Maps

You can iterate over a map using a `for` loop with the `range` keyword:

```go
package main
import "fmt"

func main() {
    m := map[string]int{"apple": 5, "banana": 3, "cherry": 8}
    for key, value := range m {
        fmt.Printf("Key: %s, Value: %d\n", key, value)
    }
}
```

**Output:** (Order may vary because maps are unordered)
```
Key: apple, Value: 5
Key: banana, Value: 3
Key: cherry, Value: 8
```

---

## 7. Nested Maps

Maps can contain other maps as values, creating a nested structure:

```go
package main
import "fmt"

func main() {
    student := map[string]map[string]string{
        "101": {
            "name":  "Alice",
            "grade": "A",
        },
        "102": {
            "name":  "Bob",
            "grade": "B",
        },
    }
    fmt.Println(student["101"]["name"]) // Output: Alice
}
```

---

## 8. Key Characteristics of Maps

- **Unordered:** The order of key-value pairs is not guaranteed.
- **Efficient:** Provides fast lookups, updates, and deletions.
- **Key Restrictions:** Keys must be of a type that supports equality comparison (e.g., strings, integers).
- **Nil Maps:** A `nil` map behaves like an empty map when reading but causes a runtime error when writing.

```go
var m map[string]int // Nil map
fmt.Println(m == nil) // Output: true
// m["key"] = 1       // This would cause a runtime error
```

---

## Summary

- Use `make()` or map literals to create maps.
- Access, update, and delete values using keys.
- Use the **comma ok** idiom to check for key existence.
- Iterate with `range`, but note that the order is not guaranteed.
- Maps are powerful for dynamic data structures where fast lookup is required.
