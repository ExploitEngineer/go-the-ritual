/*
Empty Interface

The empty interface interface{} can hold values of any type since every type implements at least zero methods. Used for generic programming before Go 1.18 generics. Requires type assertions or type switches to access underlying values. Common in APIs handling unknown data types.
*/

package main

import "fmt"

func EmptyInterfaces() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
