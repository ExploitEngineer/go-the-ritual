package compositetypes

import "fmt"

func MakeVsLiteral() {
	// What is the difference in initial values?
	// In a slice literal, the slice is initialized with actual values,
	// so both length and capacity are equal to the number of elements provided.
	//
	// Example:
	// a := []int{1, 2} → len = 2, cap = 2
	//
	// With make(), the slice is initialized with a given length,
	// and all elements are set to their zero values.
	//
	// Example:
	// b := make([]int, 2) → len = 2, cap = 2, values = [0, 0]
	//
	// The key difference is:
	// - literal: initialized with real values
	// - make(): initialized with zero values, pre-allocated size
	//
	// append() only uses capacity, not length. If capacity is exceeded,
	// Go allocates a new underlying array.
	a := []int{1, 2}
	b := make([]int, 2)

	fmt.Println(a)
	fmt.Println(b)
}
