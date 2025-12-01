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
    Name   string
    Age    int
    Gender string
}

type Police struct {
    Rank     string
    Precinct string
}

func main() {
    person := Person{Name: "John", Age: 30, Gender: "Male"}

    // Keys
    keys := obj.Keys(person)
    fmt.Println(keys) // [Name Age Gender]

    // Values
    values := obj.Values(person)
    fmt.Println(values) // [John 30 Male]

    // Entries
    entries := obj.Entries(person)
    fmt.Println(entries) // [{Name John} {Age 30} {Gender Male}]

    // Flat entries
    flat := obj.EntriesFlat(person)
    fmt.Println(flat) // [[Name John] [Age 30] [Gender Male]]

    // Pick specific fields
    picked := obj.Pick(person, "Name", "Gender")
    fmt.Println(picked["Name"], picked["Gender"]) // John Male

    // Omit specific fields
    omitted := obj.Omit(person, "Name", "Gender")
    fmt.Println(omitted["Age"]) // 30

    // Merge structs
    police := Police{Rank: "Captain", Precinct: "CA"}
    merged := obj.Merge(person, police)
    fmt.Printf("Name: %s\nRank: %s\n", merged["Name"], merged["Rank"])
    // Name: John
    // Rank: Captain

    // Convert struct to map
    policeMap := obj.ToMap(police)
    fmt.Printf("Precinct: %s\nRank: %s\n", policeMap["Precinct"], policeMap["Rank"])
    // Precinct: CA
    // Rank: Captain
}

```

### Array Utilities
// TODO, not yet done

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
// TODO, not yet done

## JavaScript/TypeScript Comparison

| JavaScript                    | Go Equivalent                 |
| ----------------------------- | ----------------------------- |
| `Object.keys(obj)`            | `obj.Keys(obj)`            |
| `Object.values(obj)`          | `obj.Values(obj)`          |
| `Object.entries(obj)`         | `obj.Entries(obj)`         |
| `Object.fromEntries(entries)` | `obj.FromEntries(entries)` |
| `_.pick(obj, keys)`           | `obj.Pick(obj, keys...)`   |
| `_.omit(obj, keys)`           | `obj.Omit(obj, keys...)`   |
| `array.map(fn)`               | `obj.Map(array, fn)`       |


## Requirements

- Go 1.18 or higher (for generics support)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see LICENSE file for details