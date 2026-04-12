package compositetypes

import "fmt"

func SliceLengthVsCapacity() {
	// explain len, cap and why extra capacity exists ?
	// length defined how many items in slice
	// capacity defines how many can be in slice in future and its 10 and if we add more values than 10 it will get new memory and make slice big without any issues
	s := make([]int, 3, 10)

	fmt.Println(len(s))
	fmt.Println(cap(s))
}
