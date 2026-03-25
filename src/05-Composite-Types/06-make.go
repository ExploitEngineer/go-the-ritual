/*
Creates and initializes slices, maps, and channels. Unlike new(), returns usable values. Examples: make([]int, 5, 10) for slices, make(map[string]int) for maps, make(chan int) for channels. Essential for initializing reference types.
*/

package main

import "fmt"

/*
When it comes to working with data structures like slices, maps, and channels, you'll likely encounter the new() and make() functions. While both are used for memory allocation, they serve distinct purposes.
*/

type Person struct {
	Name   string
	Age    uint8
	Gender string
}

func MakeNew() {
	/*
		NOTE: The new() Function
		The new() function in Go is a built-in function that allocates memory for a new zeroed value of a specified type and returns a pointer to it. It is primarily used for initializing and obtaining a pointer to a newly allocated zeroed value of a given type, usually for data types like structs.

		Here's a simple example:
	*/

	// Using new() to allocate memory for a Person struct
	p := new(Person)

	// Initializing the fields
	p.Name = "John Doe"
	p.Age = 30
	p.Gender = "Male"

	fmt.Println(p)

	/*
		In this example, new(Person) allocates memory for a new Person struct, and p is a pointer to the newly allocated zeroed value.
	*/

	/*
		NOTE: The make() Function
		On the other hand, the make() function is used for initializing slices, maps, and channels – data structures that require runtime initialization. Unlike new(), make() returns an initialized (non-zeroed) value of a specified type.

		Let's look at an example using a slice:
	*/

	// Using make() to create a slice with a specified length and capacity
	s := make([]int, 10, 15)

	// Initializing the elements
	for i := 0; i < 10; i++ {
		s[i] = i + 1
	}

	fmt.Println(s)

	/*
		In this example, make([]int, 10, 15) creates a slice of integers with a length of 10 and a capacity of 15. The make() function ensures that the slice is initialized with non-zero values.
	*/
}
