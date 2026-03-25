/*
Variadic Functions

A variadic function in Go accepts zero or more arguments of a specific type. You define them using the ... syntax before the type in the function signature. For example, func sum(nums ...int) takes any number of integers.

NOTE: Key points:
- The variadic parameter must be the last parameter in the function.
- Inside the function, the variadic parameter is treated as a slice of the specified type.
- You can pass zero, one, or many arguments to a variadic function
*/

package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

/*
Passing Slices to Variadic Functions

One cool trick is passing a slice directly to a variadic function using the ... operator. This “unpacks” the slice into individual arguments.
*/

func joinStrings(sep string, words ...string) string {
	if len(words) == 0 {
		return ""
	}
	result := words[0]
	for i := 1; i < len(words); i++ {
		result += sep + words[i]
	}
	return result
}

func VariadicFunctions() {
	fmt.Println(sum(1, 2, 3)) // Output: 6
	fmt.Println(sum(10))      // Output: 10
	fmt.Println(sum())        // Output: 0

	words := []string{"hello", "world", "golang"}
	fmt.Println(joinStrings(",", words...)) // Output: hello,world,golang
}
