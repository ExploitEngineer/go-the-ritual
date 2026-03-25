/*
error interface

Built-in interface with single Error() string method. Any type implementing this method can represent an error. Central to Go's error handling philosophy, providing consistent error representation across all Go code. Fundamental for effective error handling.
*/

package main

import "fmt"

/*
What is Error() ?

It's just a regular method with a specific name that Go expects. The error interface says: "I don't care what type you are — as long as you have a method called Error() that returns a string, you count as an error."
*/

// The built-in error interface
type error interface {
	Error() string
}

// Your custom type
type MyError struct {
	Msg string
}

// You add a method named Error() - this is what makes it satisfy the interface
func (e MyError) Error() string {
	return e.Msg
}

func ErrorInterface() {
	var err error = MyError{Msg: "something went wrong"}

	fmt.Println(err) // prints: something went wrong
}
