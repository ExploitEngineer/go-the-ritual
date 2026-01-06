// An array is a numbered sequence of elements of the same type with a fixed length. Here's an example in Go:

package main

import "fmt"

func main() {
	// declaring array
	var nums [3]int
	fmt.Println(nums) // [0 0 0]

	nums[1] = 2
	fmt.Println(nums[1])
	fmt.Println(nums) // [0 2 0]

	//  Initializing with values
	values := [3]int{1, 2, 3}
	fmt.Println(values) // [1 2 3]

	// array length
	fmt.Println(len(nums))

	// Omit the size when initializing an array
	values = [...]int{4, 5, 6}
	fmt.Println(values) // [4 5 6]

	// In Go, the length of an array is part of its type. [3]int and [4]int are considerd completely different types, even though they both hold integers (you cannot assign a [4]int to a [3]int), or even compare them directly, because their lengths don't match
	var first [3]int
	var second [4]int

	fmt.Println(first, second)
	// fmt.Println(a == b) -> compilation error

	fmt.Println(len(second))

	// length is inferred
	arr := [...]int{10, 20, 30, 40}
	fmt.Println(arr)

	/*
		 Inner Representation of Arrays

				In Go, arrays are represented as contiguous blocks of memory. This means that the elements of an array are stored one after the other in memory, making it easy to calculate the address of any element based on its index
	*/

	fmt.Println(&nums[0]) // address of first element
	fmt.Println(&nums[1]) // address of second element
	fmt.Println(&nums[2]) // address of third element

	// &nums is a pointer to the entire array (type *[3]int) when you print &nums, the fmt package recognizes it as a pointer to an array and shows the array's contents (&[1, 2, 3]) rather than a raw memory address.
	fmt.Println("nums array address: ", &nums)

	// Multi-dimensional Arrays
	var matrix [2][3]int // 2x3 matrix
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6

	fmt.Println(matrix)

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
}
