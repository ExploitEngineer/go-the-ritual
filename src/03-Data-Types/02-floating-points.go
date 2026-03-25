/*
Go supports the IEEE-754 32-bit and 64-bit floating-point numbers extensively. Let's see how to use it.
*/

package main

import "fmt"

func FloatingPoints() {
	var f float32
	var f2 float64
	f = 12.34567890123456
	f2 = 12.34567890123456
	fmt.Println(f, f2) // prints "12.345679 12.34567890123456"

	// Loss of Precision with Floating Point Numbers
	// Loss of precision will occur when 64-bit floating-point number is converted to 32-bit float.

	f = float32(f2)
	fmt.Println(f) // prints "1.2345679"
}
