/*
In Go, an array is a fixed-length collection of elements of the same type. The elements of an array are stored in contiguous memory locations. Arrays are zero-indexed, meaning the first element of an array is at index 0. The length of an array is part of its type, so arrays of different lengths are considered different types.
*/

package main

import "fmt"

func Arrays() {
	// creating an array
	var arr [4]int

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4

	// we can also declare an array using array literal syntax
	arr1 := [3]uint{200, 100, 50}

	// accessing array elements
	fmt.Println(arr1[0]) // 200
	fmt.Println(arr1[1]) // 100
	fmt.Println(arr1[2]) // 50

	// Iterating over an array in golang
	// Array elements can be accessed using a for loop. The index of the first element is 0, and the index of the last element is the length of the array minus 1.
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	// using range to iterate over elements
	for idx, v := range arr1 {
		fmt.Println(idx, ":", v)
	}

	// arrays have len function that returns the length of the array.
	fmt.Println(len(arr))
	fmt.Println(len(arr1))

	// array slicing is a technique to access a subset of an array. The syntax for array slicing is array[start:end]. The start index is inclusive, and the end index is exclusive. The start index defaults to 0, and the end index defaults to the length of the array.
	fmt.Println(arr[0:2]) // [1 2]
	fmt.Println(arr[:2])  // [1 2]
	fmt.Println(arr[2:])  // [3 4]
	fmt.Println(arr[:])   // [1 2 3 4]

	// Arrays can be compared using the == operator. The == operator returns true if the two arrays have the same length and the same elements in the same order.
	arr3 := [4]int{1, 2, 3, 4}
	arr4 := [4]int{1, 2, 3, 4}
	arr5 := [4]int{1, 2, 3, 5}
	fmt.Println(arr3 == arr4) // true
	fmt.Println(arr3 == arr5) // false
}
