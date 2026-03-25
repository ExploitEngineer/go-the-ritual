/*
Pointers with Maps & Slices

Maps and slices are reference types - passing them to functions doesn't copy underlying data. Modifications inside functions affect original. No need for explicit pointers. However, reassigning the slice/map variable itself won't affect caller unless using pointer.
*/

package main

import "fmt"

func PointersWithMapsAndSlices() {
	// Working with Maps (Pure Reference Behavior)
	myMap := map[string]int{"apples": 5}
	updateMap(myMap)
	fmt.Println("Map after update:", myMap["apples"]) // Output: 10

	// Working with Slices (The "Gotcha")
	mySlice := []int{1, 2, 3}

	// This changes the DATA but not the HEADER
	updateSliceElements(mySlice)
	fmt.Println("Slice after element update:", mySlice) // Output: [100, 2, 3]

	// This tries to APPEND but fails to reach the caller
	failToAppend(mySlice)
	fmt.Println("Slice after failed append:", mySlice) // Output: [100, 2, 3] (Still length 3!)

	// Using a Pointer to fix the Append
	reallyAppend(&mySlice)
	fmt.Println("Slice after pointer append:", mySlice) // Output: [100, 2, 3, 999]
}

func updateMap(m map[string]int) {
	m["apples"] = 10 // Changes the original map
}

func updateSliceElements(s []int) {
	s[0] = 100 // Changes the original underlying array
}

func failToAppend(s []int) {
	// This creates a NEW header locally. The caller never sees it.
	s = append(s, 999)
}

func reallyAppend(s *[]int) {
	// We dereference the pointer to modify the caller's actual header
	*s = append(*s, 999)
}
