// Go's structs are typed collections of fields. They're useful for grouping data together to form records.

package main

import "fmt"

// This person struct type has name and age fields.
type person struct {
	name string
	age  uint8
}

// newPerson constructs a new person struct with the give name.

// Go is a garbage collected language; you can safely return a pointer to a local variable - it will only be cleaned up by the garbage collector when there are no active references to it.
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	// This syntax creates a new struct
	fmt.Println(person{"Bob", 20})

	// You can name the fields when initializing a struct.
	fmt.Println(person{name: "Alice", age: 30})

	// Omitted fields will be zero-valued.
	fmt.Println(person{name: "Fred"})

	// An & perfix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40})

	// It's idiomati to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("John"))

	// Access struct fields with a dot.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// You can also use dots with struct pointers - the pointers are automatically dereferenced.e
	sp := &s
	fmt.Println(sp.age)

	// structs are mutable
	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println(s.age)

	// If a struct type is only used for a single value, we don't have to give it a name. The value can have an anonymous struct type. This technique is commonly used for table-driven-tests.
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
