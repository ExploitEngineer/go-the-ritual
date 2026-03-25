/*
Type Assertions

Extract underlying concrete value from interface. Syntax: value.(Type) or value, ok := value.(Type) for safe assertion. Panics if type assertion fails without ok form. Essential for working with interfaces and empty interfaces.
*/

package main

import "fmt"

func TypeAssertions() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
