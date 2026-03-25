/*
Methods vs Functions

Methods are functions with receiver arguments, defined outside type declaration. Enable object-like behavior on types. Functions are standalone, methods belong to specific types. Methods can have value or pointer receivers. Both can accept parameters and return values.
*/

package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

/*
FUNCTION
A standalone function. It requires the object to be passed as an argument.
*/

func AreaFunction(r Rectangle) float64 {
	return r.Width * r.Height
}

/*
THE METHOD (value receiver)
This is "attached" to the Rectangle type.
It receives a COPY of the rectangle.
*/

func (r Rectangle) AreaMethod() float64 {
	return r.Width * r.Height
}

/*
THE METHOD (Pointer Receiver)
Use a pointer receiver if you need to MODIFY the original object
or avoid copying a large struct.
*/

func (r *Rectangle) Scale(factor float64) {
	r.Width = r.Width * factor
	r.Height = r.Height * factor
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}

	// Calling the Function
	fmt.Println("Area (Function):", AreaFunction(rect))

	// Calling the Method
	fmt.Println("Area (Method):", rect.AreaMethod())

	// Modifying via Pointer Method
	rect.Scale(2)
	fmt.Printf("Scaled Rect: %+v\n", rect)
}
