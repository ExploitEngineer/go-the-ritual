/*
Convert values between different types using `Type(value)` syntax. Go requires explicit conversion even between related types like `int` and `int64`. Essential for working with different data types and ensuring type compatibility in programs.
*/

package main

import "fmt"

func TypeConversion() {
	/*
		Numeric Types
		Conversion between different numeric types are straightforward but must be explicit.
	*/
	i := 42
	f := float64(i) // int to float64
	u := uint(i)    // int to uint

	fmt.Println(i, f, u)

	/*
		String and Byte Slice
		Go strings are immutable, but they can be converted to and from byte slices ([]byte), which are mutable
	*/
	s := "hello"
	b := []byte(s)  // string to []byte
	s2 := string(b) // []byte to string

	fmt.Printf("%s %b %s\n", s, b, s2)

	// Similarly, you can convert between strings and rune slices ([]rune), where rune is a type alias for int32
	r := []rune(s)  // string to []rune
	s3 := string(r) // []rune to string

	fmt.Println(r, s3)

	/*
		Custom Type Conversion
		In Go, you can define youw own types based on existing ones. Conversions between custom types and their underlying types are explicit.
	*/
	type MyInt int
	i1 := 10
	mi := MyInt(i1) // int to MyInt
	i2 := int(mi)   // MyInt to int

	fmt.Println(i1, mi, i2)

	/*
		Pointer Conversion
		Pointers in Go reference the memory address of a variable. You can convert between a value and its pointer.
	*/
	x := 42
	p := &x // int to *int (pointer to int)
	y := *p // *int to int (dereferencing)

	fmt.Println(x, p, y)
}
