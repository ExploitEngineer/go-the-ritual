/*
Slices in Go:

- Slices are dynamic-size sequences built on top of arrays.
- They do not store data themselves; instead, they reference an underlying array.
- Slices can grow or shrink using operations like append.
- Syntax looks similar to arrays, but slices do not include a fixed size.

A slice keeps track of three things:

1. Pointer: A reference to the underlying array where the data is stored.
2. Length: The number of elements currently in the slice.
3. Capacity: The maximum number of elements the slice can hold before a new underlying array is allocated (always ≥ length).

Key points:

- Unlike arrays, slices are typed only by the element type, not by the number of elements.
- An uninitialized slice has the value `nil` and a length of 0.
- Slices are reference types, so multiple slices can share the same underlying array.
- You can create slices in three main ways:
  1. From an existing array: `s := arr[1:4]`
  2. Using slice literals: `s := []int{1, 2, 3}`
  3. Using `make()`: `s := make([]int, length, capacity)`

*/

package main

import "fmt"

func main() {
	// creating slice using an existing array
	arr := [5]uint8{1, 2, 3, 4, 5}
	s := arr[1:4]  // slice from index 1 to 3 (4 is excluded)
	fmt.Println(s) // [2 3 4]

	// Using slice literals
	s1 := []uint8{10, 20, 30}
	fmt.Println(s1)

	/*
				Using make

		Syntax: make([]T, length, capacity)
		- length -> how many elements are initially in the slice.
		- capacity -> underlying array size. Optional; defaults to length if not provided
	*/
	s2 := make([]int, 5) // slice of length 5, cpacity 5, initialized to 0
	fmt.Println(s2)

	s3 := make([]uint8, 3, 5)
	fmt.Println(s3)      // [0 0 0] length = 3, capacity = 5
	fmt.Println(len(s3)) // 3
	fmt.Println(cap(s3)) // 5

	s3 = append(s3, 10, 20)
	fmt.Println(s3) // [0 0 0 10 20]

	// Slices are references
	array := [3]int{1, 2, 3}
	slice1 := array[:]
	slice2 := slice1[:2]
	slice2[0] = 100
	fmt.Println(array)  // [100 2 3]
	fmt.Println(slice1) // [100 2 3]
	fmt.Println(slice2) // [100 2]
}
