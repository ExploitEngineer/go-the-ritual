/*
Error Handling Basics

Go uses explicit error handling with error return values. Functions return error as last value. Check if err != nil pattern. Create errors with errors.New() or fmt.Errorf(). No exceptions - errors are values to be handled explicitly.
*/

package main

import (
	"errors"
	"fmt"
)

/*
Using the New function

Golang error package has a function called New() which can be used to create errors easily.
*/

func e(v int) (int, error) {
	if v == 0 {
		return 0, errors.New("Zero cannot be used.")
	} else {
		return 2 * v, nil
	}
}

func ErrorHandlingBasics() {
	v, err := e(0)
	if err != nil {
		fmt.Println(err, v)
	}
}
