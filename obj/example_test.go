package obj_test

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

func ExampleKeys() {
	person := Person{Name: "John", Age: 30, Gender: "Male"}
	keys := obj.Keys(person)
	fmt.Println(keys)
	// Output: [Name Age Gender]
}

func ExampleValues() {
	person := Person{Name: "John", Age: 30, Gender: "Male"}
	values := obj.Values(person)
	fmt.Println(values)
	// Output: [John 30 Male]
}

func ExampleEntries() {
	person := Person{Name: "John", Age: 30, Gender: "Male"}
	entries := obj.Entries(person)
	fmt.Println(entries)
	// Output: [{Name John} {Age 30} {Gender Male}]
}

func ExampleEntriesFlat() {
	person := Person{Name: "John", Age: 30, Gender: "Male"}
	entries := obj.EntriesFlat(person)
	fmt.Println(entries)
	// Output: [[Name John] [Age 30] [Gender Male]]
}

func ExamplePick() {
	p := Person{
		Name:   "John",
		Age:    30,
		Gender: "Male",
	}

	picked := obj.Pick(p, "Name", "Gender")
	fmt.Println(picked["Name"], picked["Gender"])
	// Output: John Male
}

func ExampleOmit() {
	p := Person{
		Name:   "John",
		Age:    30,
		Gender: "Male",
	}

	omitted := obj.Omit(p, "Name", "Gender")
	fmt.Println(omitted["Age"])
	// Output: 30
}

func ExampleMerge() {
	person := Person{
		Name:   "John",
		Age:    30,
		Gender: "Male",
	}

	police := Police{
		Rank:     "Captain",
		Precinct: "CA",
	}

	merged := obj.Merge(person, police)

	fmt.Printf("Name: %s\nRank: %s\n", merged["Name"], merged["Rank"])
	// Output:
	// Name: John
	// Rank: Captain

}

func ExampleToMap() {
	p := Police{
		Rank:     "Captain",
		Precinct: "CA",
	}

	police := obj.ToMap(p)

	fmt.Printf("Precinct: %s\nRank: %s\n", police["Precinct"], police["Rank"])
	// Output:
	// Precinct: CA
	// Rank: Captain

}
