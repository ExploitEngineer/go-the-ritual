/*
Function Basics

Reusable code blocks declared with func keyword. Support parameters, return values, multiple returns. First-class citizens - can be assigned to variables, passed as arguments. Fundamental building blocks for organizing code logic.
*/

package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func FunctionBasics() {
	fmt.Println(add(42, 13))
}
