/*
Generic Types / Interfaces

Create reusable data structures and interface definitions working with multiple types. Define with type parameters like type Container[T any] struct { value T }. Enable type-safe containers, generic slices, maps, and custom structures while maintaining Go's strong typing.
*/

package main

import "fmt"

/*
A Generic Interface
This interface says: "I work with any type T, as long as T can be 'Processed'."
*/

type Processor[T any] interface {
	Process(data T) string
}

/*
A Generic Struct (The Container)
 This struct can hold a value of any type T.
*/

type Container[T any] struct {
	Value T
}

// Implementing the Interface for specific types

type StringProcessor struct{}

func (sp StringProcessor) Process(data string) string {
	return "Processing string: " + data
}

type NumberProcessor struct{}

func (np NumberProcessor) Process(data int) string {
	return fmt.Sprintf("Processing number: %d", data*2)
}

// A Generic Function that uses both

func Execute[T any](p Processor[T], val T) {
	fmt.Println(p.Process(val))
}

func GenericTypesInterfaces() {
	// Using the StringProcessor with a string
	sp := StringProcessor{}
	Execute[string](sp, "Hello Go!")

	// Using the NumberProcessor with an int
	np := NumberProcessor{}
	Execute[int](np, 42)

	// Using the Generic Container
	intBox := Container[int]{Value: 100}
	strBox := Container[string]{Value: "Secret Message"}

	fmt.Printf("Container 1: %v, Container 2: %v\n", intBox.Value, strBox.Value)
}
