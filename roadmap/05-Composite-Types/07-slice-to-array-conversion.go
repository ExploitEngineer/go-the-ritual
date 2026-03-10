/*
Slice to Array Conversion
Convert slice to array using [N]T(slice) (Go 1.17+). Copies data from slice to fixed-size array. Panics if slice has fewer than N elements. Useful when array semantics or specific size guarantees are needed.
*/

package main

import "fmt"

func SliceToArray() {
	// Create a slice
	s := []int{10, 20, 30, 40, 50}

	// Convert slice to array of size 3
	arr := [3]int(s)

	fmt.Println("Slice:", s)
	fmt.Println("Array:", arr)
}
