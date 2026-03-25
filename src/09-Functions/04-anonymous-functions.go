/*
Anonymous Functions

Functions declared without names, also called function literals or lambdas. Can be assigned to variables, passed as arguments, or executed immediately. Useful for short operations, callbacks, goroutines, and closures. Access enclosing scope variables. Common in event handlers and functional patterns.
*/

package main

import "fmt"

func AnonymousFunctions() {
	// Declaring Anonymous Function
	func() {
		fmt.Println("Inside anonymous function")
	}()

	// Assigning Anonymous Function to a Variable
	v := func() {
		fmt.Println("Inside anonymous function")
	}
	v()

	// Passing Arguments to an Anonymous Function
	func(v int) {
		fmt.Println(v)
	}(10)
}
