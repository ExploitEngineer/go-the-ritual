/*
var declares 1 or more variables.

You can declare multiple variables at once.
*/

package main

import "fmt"

func main() {
	// String
	var a string = "initial"
	fmt.Println(a)

	// Integers
	var b, c int = 1, 2
	fmt.Println(b, c)

	var i8 int8 = 127
	var u8 uint8 = 255
	fmt.Println(i8, u8)

	var i16 int16 = 32767
	var u16 uint16 = 65535
	fmt.Println(i16, u16)

	var i32 int32 = 2147483647
	var u32 uint32 = 4294967295
	fmt.Println(i32, u32)

	var i64 int64 = 9223372036854775807
	var u64 uint64 = 18446744073709551615
	fmt.Println(i64, u64)

	// Boolean
	var d bool = true
	fmt.Println(d)

	// Float
	var f float32 = 3.14
	var g float64 = 3.1415926535
	fmt.Println(f, g)

	// Complex
	var c1 complex64 = complex(1, 2)
	var c2 complex128 = complex(1.123456789, 2.987654321)

	// Go will inter the type of initialized variables
	apple := "apple"
	fmt.Println(apple)

	fmt.Println(c1, c2)
}
