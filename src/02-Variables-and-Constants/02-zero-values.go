/*
Default values for uninitialized variables: `0` for numbers, `false` for booleans, `""` for strings, `nil` for pointers/slices/maps, Ensures predictable initial state and reduces initialization errors. Fundamental for reliable Go code.
*/

package main

import "fmt"

func zeroValues() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
