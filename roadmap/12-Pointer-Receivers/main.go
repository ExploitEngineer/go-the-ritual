/*
Pointer Receivers

Methods receive pointer to struct rather than copy using func (p *Type) methodName() syntax. Necessary when method modifies receiver state or struct is large. Go automatically handles value/pointer conversion when calling methods.
*/

package main

import "fmt"

type Coordinates struct {
	X, Y int
}

// Pointer Receiver: Receives the ADDRESS of 'c'
// This allows the method to modify the original fields directly.

func (c *Coordinates) Move(newX, newY int) {
	c.X = newX
	c.Y = newY
}

func main() {
	point := Coordinates{X: 10, Y: 10}
	fmt.Printf("Starting Position: %+v\n", point)

	// We call the method on 'point'
	// Go automatically turns this into (&point).Move(...) for us!
	point.Move(50, 80)

	// The original 'point' has been updated
	fmt.Printf("New Position: %+v\n", point)
}
