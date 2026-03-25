/*
Generic Functions

Write functions working with multiple types using type parameters in square brackets like func FunctionName[T any](param T) T. Enable reusable algorithms maintaining type safety. Particularly useful for utility functions and data processing that don't depend on specific types.
*/

package main

import "fmt"

/*
T is the "Type Parameter"
any is a constraint meaning T can be literally anything
*/

func PrintSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func GenericFunctions() {
	// works with a slice of strings
	names := []string{"Alice", "Bob"}
	PrintSlice(names)

	// works with the slice of integers
	numbers := []uint8{10, 20, 30}
	PrintSlice(numbers)
}
