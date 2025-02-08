# Structs in Go (Golang) - Comprehensive Guide

A **struct** in Go is a user-defined type that groups related fields together. It is similar to classes in object-oriented languages but does not support inheritance.

---

## 1. Declaring a Struct

A struct is defined using the `type` keyword.

```go
package main
import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    var p Person // Declaring a struct variable
    fmt.Println(p) // Output: { "" 0 }
}
```

---

## 2. Initializing a Struct

### 2.1 Using Struct Literal

```go
package main
import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{Name: "John", Age: 30} // Initializing struct
    fmt.Println(p) // Output: {John 30}
}
```

### 2.2 Initializing Without Field Names

```go
p := Person{"Alice", 25} // Fields are assigned in order
```

**Note:** This method is not recommended as it reduces code readability.

---

## 3. Accessing and Modifying Struct Fields

```go
package main
import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{Name: "Alice", Age: 25}
    fmt.Println("Before:", p.Name, p.Age)

    p.Age = 26 // Modifying struct field
    fmt.Println("After:", p.Name, p.Age)
}
```

**Output:**
```
Before: Alice 25
After: Alice 26
```

---

## 4. Struct with Methods

Structs can have methods by defining functions with a receiver.

```go
package main
import "fmt"

type Person struct {
    Name string
    Age  int
}

// Method with value receiver
func (p Person) Greet() {
    fmt.Println("Hello, my name is", p.Name)
}

func main() {
    p := Person{Name: "John", Age: 30}
    p.Greet()
}
```

**Output:**
```
Hello, my name is John
```

---

## 5. Struct with Pointer Receiver (Modifying Fields in Methods)

```go
package main
import "fmt"

type Person struct {
    Name string
    Age  int
}

// Method with pointer receiver to modify struct
func (p *Person) IncrementAge() {
    p.Age++
}

func main() {
    p := Person{Name: "Alice", Age: 25}
    p.IncrementAge()
    fmt.Println("New Age:", p.Age) // Output: New Age: 26
}
```

---

## 6. Anonymous Structs

You can define a struct without explicitly declaring a type.

```go
package main
import "fmt"

func main() {
    p := struct {
        Name string
        Age  int
    }{Name: "Bob", Age: 40}

    fmt.Println(p)
}
```

---

## 7. Nested Structs

Structs can contain other structs as fields.

```go
package main
import "fmt"

type Address struct {
    City  string
    State string
}

type Person struct {
    Name    string
    Age     int
    Address Address // Nested struct
}

func main() {
    p := Person{Name: "Alice", Age: 28, Address: Address{City: "New York", State: "NY"}}
    fmt.Println(p)
    fmt.Println("City:", p.Address.City)
}
```

---

## Summary

- **Structs** group related fields together.
- **Methods** can be defined for structs using value or pointer receivers.
- **Anonymous structs** are useful for temporary data structures.
- **Nested structs** help in creating complex data models.

Structs in Go provide a powerful way to structure and manipulate data efficiently.
