/*
Wrapping/Unwrapping Errors

Create error chains preserving original errors while adding context using fmt.Errorf() with %w verb. Use errors.Unwrap(), errors.Is(), and errors.As() to work with wrapped errors. Enables rich error contexts for easier debugging.
*/
package main

import (
	"errors"
	"fmt"
)

/*
errors.Unwrap
Peels off one layer of wrapping and returns the inner error.
*/

var ErrNotFound = errors.New("user not found")

func unwrapExample() {
	wrapped := fmt.Errorf("database failed: %w", ErrNotFound)

	fmt.Println(wrapped)                // database failed: user not found
	fmt.Println(errors.Unwrap(wrapped)) // user not found
}

/*
errors.Is
Checks if a specific error exists anywhere in the chain, no matter how deep.
*/

func findCustomer(id int) error {
	return fmt.Errorf("findCustomer: %w",
		fmt.Errorf("db query: %w", ErrNotFound)) // two layers deep
}

func isExample() {
	err := findCustomer(42)

	// digs through all layers to find ErrNotFound
	if errors.Is(err, ErrNotFound) {
		fmt.Println("it is a not found error")
	}
}

/*
errors.As
Like errors.Is but extracts the concrete error type so you can read its fields.
*/

type ValidationError struct {
	Field string
}

func (e *ValidationError) Error() string {
	return "validation failed on field: " + e.Field
}

func process() error {
	return fmt.Errorf("process failed: %w", &ValidationError{Field: "email"})
}

func asExample() {
	err := process()

	var valErr *ValidationError
	if errors.As(err, &valErr) {
		fmt.Println("bad field:", valErr.Field) // bad field: email
	}
}

/*
Full chain example
Shows how errors travel up through functions, each layer adding context.
*/

type DBError struct {
	Code int
}

func (e *DBError) Error() string {
	return fmt.Sprintf("db error code %d", e.Code)
}

func queryDB() error {
	return &DBError{Code: 404}
}

func findCustomerFull(id int) error {
	err := queryDB()
	return fmt.Errorf("findCustomer id %d: %w", id, err)
}

func fullChainExample() {
	err := findCustomerFull(42)

	// full message with all context
	fmt.Println(err) // findCustomer id 42: db error code 404

	// dig in for the specific type and read its fields
	var dbErr *DBError
	if errors.As(err, &dbErr) {
		fmt.Println("db code was:", dbErr.Code) // db code was: 404
	}
}

func WrappingUnwrappingErrors() {
	fmt.Println("--- errors.Unwrap ---")
	unwrapExample()

	fmt.Println("\n--- errors.Is ---")
	isExample()

	fmt.Println("\n--- errors.As ---")
	asExample()

	fmt.Println("\n--- Full chain ---")
	fullChainExample()
}
