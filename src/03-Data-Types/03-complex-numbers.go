/*
There are two complex types in Go. The complex64 and the complex128.
*/

package main

import "fmt"

func ComplexNumbers() {
	// Initializing Complex Numbers
	// Initializing complex numbers is really easy. You can use the constructor or the initialization syntax as well.
	c1 := complex(10, 11) // constructor init
	c2 := 10 + 11i        // complex number init syntax

	fmt.Println(c1, c2)

	// Parts of a Copmlex Number
	// There are two parts of a complex number. The real and imaginary part. We use functions to get those.
	c := complex(23, 31)
	realPart := real(c) // gets real part
	imagPart := imag(c) // gets imaginary part

	fmt.Println(realPart, imagPart)

	// Operations on a Complex Number
	// A complex variable can do any operation like addition, subtraction, multiplication, and division.

	cmp1 := complex(2, 3)
	cmp2 := 4 * 5i
	cmp3 := cmp1 + cmp2
	fmt.Println("Add: ", cmp3)

	re := real(cmp3)
	im := imag(cmp3)
	fmt.Println(re, im)
}
