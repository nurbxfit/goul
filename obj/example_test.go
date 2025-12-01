package obj_test

import (
	"fmt"

	"github.com/nurbxfit/goul/obj"
)

type Person struct {
	Name string
	Age  int
}

func ExampleKeys() {
	person := Person{Name: "John", Age: 30}
	keys := obj.Keys(person)
	fmt.Println(keys)
	// Output: [Name Age]
}

func ExampleValues() {
	person := Person{Name: "John", Age: 30}
	values := obj.Values(person)
	fmt.Println(values)
	// Output: [John 30]
}

func ExampleEntries() {
	person := Person{Name: "John", Age: 30}
	entries := obj.Entries(person)
	fmt.Println(entries)
	// Output: [{Name John} {Age 30}]
}
