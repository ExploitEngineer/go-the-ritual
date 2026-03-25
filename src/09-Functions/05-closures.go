/*
Closures

Functions capturing variables from surrounding scope, accessible even after outer function returns. "Close over" external variables for specialized functions, callbacks, state maintenance. Useful for event handling, iterators, functional programming. Important for flexible, reusable code.
*/

package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// another example of clousre
func createCounter() func() int {
	count := 0
	increment := func() int {
		count++
		return count
	}
	return increment
}

func Closures() {
	pos, neg := adder(), adder()
	for i := range 10 {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	counter1 := createCounter()
	counter2 := createCounter()

	fmt.Println(counter1()) // Output: 1
	fmt.Println(counter1()) // Output: 2
	fmt.Println(counter2()) // Output: 1
	fmt.Println(counter2()) // Output: 2
}
