// A variadic function ia a function that can accept zero or move values of the same type as its last parameter

package main

import "fmt"

// nums is a slice of int: []int
func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	fmt.Println("slice:", nums)

	sum(nums...)
}
