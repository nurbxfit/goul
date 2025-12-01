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
