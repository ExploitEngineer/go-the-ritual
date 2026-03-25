/*
errors.New

Simplest way to create error values by taking a string message and returning an error implementing the error interface. Useful for simple, static error messages. Often combined with error wrapping or used for predefined error constants.
*/

package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func ErrorsNew() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
