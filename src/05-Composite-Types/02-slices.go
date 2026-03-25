/*
In Go, a slice is a dynamic, flexible view into an underlying array. Unlike arrays, slices can grow and shrink dynamically, making them more versatile and commonly used in Go programming.

Key Characteristics
- Slices are reference types
- Modifying a slice can effect the underlying array

A slice consist of three key components:
- Pointer to the underlying array
- Length of the slice
- Capacity of the slice
*/

package main

import "fmt"

func Slices() {
	// using slice literal to create slice
	fruits := []string{"apple", "banana", "orange"}
	fmt.Println(fruits)

	// using make() function
	/*
		when using make function for slices we can provide three parameters.
		- type
		- length
		- capacity
	*/
	numbers := make([]int, 5)        // Length 5, capacity 5
	newNumbers := make([]int, 5, 10) // Length 5, capacity 10

	fmt.Println(numbers)
	fmt.Println(newNumbers)

	// creating slice from an existing array
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4] // creates a slice from index 1 to 3

	fmt.Println(slice)

	// appending elements
	s1 := []string{"hi", "how", "are"}
	resultSlice := append(s1, "you ?")
	fmt.Println(resultSlice)

	// slice indexing and slicing
	numbers = []int{0, 1, 2, 3, 4, 5}
	subset := numbers[2:4]

	fmt.Println(numbers, subset)
}
