/*
Array to Slice Conversion
Convert arrays to slices using expressions like array[:] or array[start:end]. Creates slice header pointing to array memory - no data copying. Modifications through slice affect original array. Efficient way to use arrays with slice-based APIs.
*/

package main

import "fmt"

func ArrayToSlice() {
	arr := [5]int{10, 20, 30, 40, 50}

	// Create slice from array
	s := arr[1:4]

	fmt.Println("Before change:")
	fmt.Println("Array:", arr)
	fmt.Println("Slice:", s)

	// Modify slice
	s[0] = 999

	fmt.Println("\nAfter change:")
	fmt.Println("Array:", arr)
	fmt.Println("Slice:", s)
}
