/*
Signed integers (int8, int16, int32, int64) handle positive/negative numbers. Unsigned (uint8, uint16, uint32, uint64) handle only non-negative but larger positive range. `int`/`uint` are platform-dependent. Choose based on range and memory needs.

*Signed Integers*
Signed integers in Go can represent both positive and negative numbers. Go provides several signed integer types with different sizes:

- int8 - Represents an 8-bit signed integer. Range: -128 to 127.
- int16 - Represents a 16-bit signed integer. Range: -32,768 to 32,767
- int32 — Represents a 32-bit signed integer. Range: -2,147,483,648 to 2,147,483,647
- int64 — Represents a 64-bit signed integer. Range: -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807.

*Unsigned Integers*
Unsigned integers, on the other hand, can only represent non-negative numbers. The unsigned integer type in Go include:

- uint8 - Represents an 8-bit unsinged integer. Range: 0 to 255.
- uint16 —Represents a 16-bit unsigned integer. Range: 0 to 65,535.
- uint32 —Represents a 32-bit unsigned integer. Range: 0 to 4,294,967,295.
- uint64 —Represents a 64-bit unsigned integer. Range: 0 to 18,446,744,073,709,551,615

Besides these, Go also offers int and uint, which are platform-dependent. They are 32 bits in size on 32-bit systems and 64 bits in size on 64-bit systems. There’s also the uintptr type, an unsigned integer large enough to store the uninterpreted bits of a pointer value, but we’ll save that story for another day.
*/

package main

import "fmt"

func main() {
	var x uint8
	fmt.Println("Throws integer overflow", x)
	// x = 267 // range of uint8 is 0-255

	var a uint8 = 255
	fmt.Println("Unsigned 8 bit integer", a)

	var b int8 = -127
	fmt.Println("Signed 8 bit integer,", b)

	// Type conversion in Golang
	var w uint32
	var y uint32
	var z uint8
	fmt.Println("Type Conversion")
	w = 26700
	y = uint32(x)        // data preserved number is inside range
	z = uint8(x)         // data loss due to out of range conversion
	fmt.Println(w, y, z) // prints 26700 26700 76
}
