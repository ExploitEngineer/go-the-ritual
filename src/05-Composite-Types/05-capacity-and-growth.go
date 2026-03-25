/*
What is a Slice in Go?
Think of a slice as a resizable box that holds a list of items. You can start with just a few items, but if you need more room, the box can automically grow bigger to fit more stuff.

A slice has three main properties:

1. Length: How many items are currently in the slice?
2. Capacity: How many items can the slice hold before needing to grow?
3. Pointer: A reference to the underlying "storage area" (called an array) where the items are kept.

Slice capacity determines when reallocation occurs during append operations. Go typically doubles capacity for smaller slices. Pre-allocate with make([]T, length, capacity) to optimize memory usage and minimize allocations in performance-critical code.
*/

package main

import "fmt"

func CapacityAndGrowth() {
	/*
	   How Do Slices Grow?
	   When you create a slice, you can either specify its size upfront or leave it to grow dyanmically. For example:
	*/

	nums := []int{1, 2, 3} // Slices with 3 items
	fmt.Println(nums)

	/*
		Here the slice has:

		- Length: 3 (because it holds 3 items)
		- Capacity: 3 (because the initial storage space matches the length)

		If you try to add more items using append:
	*/

	nums = append(nums, 4) // Add another item
	fmt.Println(nums)

	/*
		Go checks if the slice has enough capacity. It not:

		1.  It creates a bigger storage area (usually double the size)
		2. It copies the existing items into the new storage
		3. It adds the new item



		Small Slices (Rapid Growth)
		For smaller slices, the capacity typically doubles when you exceed it. This ensures that you don't keep reallocating for every small addition.

		Example Growth:
		Capacity: 2 -> 4 -> 8 -> 16 -> ....

		Large Slices (Slower Growth)
		As slices get larger, the growth becomes more conservative to save memory. Instead of doubling, Go increases capacity by about 25%. For intance:

		Capacity: 1024 -> 1280 -> 1600 -> ....

	*/
}
