# GoUL

A Go utility library inspired by JavaScript/TypeScript, bringing familiar patterns to Go development.

## Installation

```bash
go get github.com/nurbxfit/goul
```

## Usage

### Object Utilities

```go
package main

import (
    "fmt"
    "github.com/nurbxfit/goul/obj"
)

type Person struct {
    Name  string
    Age   int
    Email string
}

func main() {
    person := Person{
        Name:  "John Doe",
        Age:   30,
        Email: "john@example.com",
    }

    // Get keys (like Object.keys() in JS)
    keys := obj.Keys(person)
    fmt.Println(keys) // [Name Age Email]

    // Get values (like obj.values() in JS)
    values := obj.Values(person)
    fmt.Println(values) // [John Doe 30 john@example.com]

    // Get entries (like obj.entries() in JS)
    entries := obj.Entries(person)
    for _, entry := range entries {
        fmt.Printf("%s: %v\n", entry.Key, entry.Value)
    }

    // Pick specific fields
    picked := obj.Pick(person, "Name", "Email")
    fmt.Println(picked) // map[Name:John Doe Email:john@example.com]

    // Omit specific fields
    omitted := obj.Omit(person, "Age")
    fmt.Println(omitted) // map[Name:John Doe Email:john@example.com]
}
```

### Array Utilities

```go
package main

import (
    "fmt"
    "github.com/nurbxfit/goul/arr"
)

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    
    // Map function (like Array.map() in JS)
    doubled := arr.Map(numbers, func(n int) int {
        return n * 2
    })
    fmt.Println(doubled) // [2 4 6 8 10]

    // Transform types
    strings := arr.Map(numbers, func(n int) string {
        return fmt.Sprintf("Number: %d", n)
    })
    fmt.Println(strings) // [Number: 1 Number: 2 ...]
}
```

## API Reference

### object.Keys(s interface{}) []string
Returns field names of a struct as a slice of strings.

### object.Values(s interface{}) []interface{}
Returns field values of a struct as a slice.

### object.Entries(s interface{}) []Entry
Returns key-value pairs as Entry structs.

### object.EntriesFlat(s interface{}) [][2]interface{}
Returns key-value pairs as 2D array (JS-style).

### object.FromEntries(entries []Entry) map[string]interface{}
Converts entries back to a map.

### object.Pick(obj interface{}, keys ...string) map[string]interface{}
Creates a map with only specified keys.

### object.Omit(obj interface{}, keys ...string) map[string]interface{}
Creates a map excluding specified keys.

### object.Map[T any, U any](input []T, f func(T) U) []U
Transforms a slice using a mapping function (generic).

## JavaScript/TypeScript Comparison

| JavaScript | Go Equivalent |
|------------|---------------|
| `Object.keys(obj)` | `object.Keys(obj)` |
| `Object.values(obj)` | `object.Values(obj)` |
| `Object.entries(obj)` | `object.Entries(obj)` |
| `Object.fromEntries(entries)` | `object.FromEntries(entries)` |
| `_.pick(obj, keys)` | `object.Pick(obj, keys...)` |
| `_.omit(obj, keys)` | `object.Omit(obj, keys...)` |
| `array.map(fn)` | `object.Map(array, fn)` |

## Requirements

- Go 1.18 or higher (for generics support)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see LICENSE file for details