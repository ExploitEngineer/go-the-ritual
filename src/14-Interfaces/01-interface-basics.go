/*
Interfaces Basics

Define contracts through method signatures. Types automatically satisfy interfaces by implementing required methods. Declared with type InterfaceName interface{} syntax. Enable polymorphism and flexible, testable code depending on behavior rather than concrete types.
*/

package main

import "fmt"

// Speaker defines the contract
type Speaker interface {
	Speak() string
}

// Dog is a concrete type
type Dog struct {
	Name string
}

// Dog implements Speaker implicitly

func (d Dog) Speak() string {
	return "Woof!"
}

// Robot is another concrete type
type Robot struct {
	ID string
}

// Robot also implements Speaker implicitly

func (r Robot) Speak() string {
	return "Beep Boop."
}

// This function doesn't care WHAT you are, only that you can Speak()

func DistributeMessage(s Speaker) {
	fmt.Println("This entity says:", s.Speak())
}

func InterfaceBasics() {
	d := Dog{Name: "Buddy"}
	r := Robot{ID: "RX-9"}

	// Both satisfy the Speaker interface
	DistributeMessage(d)
	DistributeMessage(r)
}
