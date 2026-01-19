/*
Constants declared with `const` represent unchanging compile-time values. `iota` creates successive integer constants starting from zero, resetting per `const` block. Useful for enumerations, bit flags, and constant sequences without manual values.


Constants in Go are immutable values that are known at compile time.
*/

package main

import "time"

type ByteSize float64

func constAndIota() {
	// Single constant
	const Pi = 3.14159

	// Multiple constants
	const (
		StatusOk    = 200
		StatusError = 500
	)

	// Typed constants
	const (
		Timeout time.Duration = 30 * time.Second
		MaxSize int           = 1024
		Debug   bool          = false
	)

	// Untyped constants
	const (
		PI    = 3.14159
		Hello = "Hello, World"
		Yes   = true
	)

	// Using iota

	// Advanced iota patterns
	const (
		// Bit shifting
		_  = iota // ignore the first value by assigning to blank identifier
		KB = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)

	// Basic iota
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	const (
		// Skip values
		_          = iota
		DebugLevel = 1 << iota
		Info
		Warning
		Error
	)

	const (
		// Offset and multiplier
		Offset = 2*iota + 1
		_
		Value
		_
		Result
	)
}
