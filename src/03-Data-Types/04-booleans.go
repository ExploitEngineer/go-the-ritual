/*
Bollean values are those with can be assigned true or false and has the type bool with it.
*/

package main

import "fmt"

func Booleans() {
	// In this code "bVal" is not initialized and thus has a zero-value, The zero-value for a boolean is false.
	var bVal bool // default is false
	fmt.Printf("bVal :%v\n", bVal)

	// Boolean operators in Go
	// Boolean operators are those operators that compare one value to others.
	a := 3
	b := 4
	fmt.Println(a == b) // false
	fmt.Println(a != b) // true
	fmt.Println(a < b)  // true
	fmt.Println(a > b)  // false
	fmt.Println(a >= b) // false
	fmt.Println(a <= b) // true
}
