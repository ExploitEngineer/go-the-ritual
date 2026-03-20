/*
Type Constraints

Specify which types can be used as type arguments for generics. Defined using interfaces with method signatures or type sets. Common constraints include any, comparable, and custom constraints. Enable writing generic code that safely operates on type parameters.
*/

package main

func max[T int | uint](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func TypeConstraints() {
	max(int(1), int(2))
	max(uint(1), uint(5))
}
