/*
Value Receivers

Methods receive copy of struct rather than pointer. Use func (v Type) methodName() syntax. Appropriate when method doesn't modify receiver or struct is small. Can be called on both values and pointers with Go automatically dereferencing.
*/

package main

import "fmt"

type Player struct {
	Name  string
	Score int
}

// Value Receiver: Receives a COPY of 'p'
// Even if we change 'p.Score' here, the original remains 0.

func (p Player) GetGreeting() string {
	p.Score = 999 // This change only exists inside this function's scope
	return fmt.Sprintf("Hello %s! Your temporary score is %d.", p.Name, p.Score)
}

func main() {
	p1 := Player{Name: "Alice", Score: 0}

	fmt.Println(p1.GetGreeting())

	// The original score is still 0 because the method used a copy
	fmt.Printf("Original Score for %s: %d\n", p1.Name, p1.Score)
}
