// In Go, an array is a numbered sequence of elements of a specific length. In typical Go code, slices are much more common; arrays are useful in some special scenarios.

// NOTE: Array types are non-dimensional, but you can compose types to build multi-dimensional data structures.

package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	// here length is defined
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// here length is inferred
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// NOTE: Initialize only specific elements
	// 3:400 means: assings 400 to array index 3 (fourth element)
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	var twoD [2][3]int

	fmt.Println("two D array:", twoD)

	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)

	// Composite literal
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)

	// NOTE: Inner Representation of arrays
	// In Go, arrays are represented as contiguous blocks of memory. This means that the elements of an array are stored one after the other in memory, making it easy to calculate the address of any element based on its index.

	nums := [3]int32{1, 2, 3} // array of 3 32-bit integers
	fmt.Println(&nums[0])
	fmt.Println(&nums[1])
	fmt.Println(&nums[2])
}
