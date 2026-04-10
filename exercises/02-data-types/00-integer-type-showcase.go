// Package datatypes have all the practice exercises about data types in go
package datatypes

import "fmt"

func IntegerTypeShowcase() {
	var i8 int8 = 127
	var i16 int16 = 255
	var ui8 uint8 = 255
	var ui16 uint16 = 355

	fmt.Println(i8, i16, ui8, ui16)
}
