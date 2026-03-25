/*
fmt.Errorf

Creates formatted error messages using printf-style verbs. Supports %w verb for error wrapping (Go 1.13+) to create error chains preserving original errors while adding context. Essential for descriptive errors with dynamic values and debugging information.
*/

package main

import "fmt"

func findUser(id int) error {
	if id <= 0 {
		return fmt.Errorf("user id %d is not vaild", id)
	}
	return nil
}

func FmtErrorf() {
	err := findUser(-1)
	if err != nil {
		fmt.Println(err)
	}
}
