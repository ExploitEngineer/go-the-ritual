/*
Type Switch

Special form of switch statement that operates on types rather than values. Syntax: switch v := i.(type). Used with interfaces to determine underlying concrete type. Each case specifies types to match. Essential for handling interface{} and polymorphic code.
*/

package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n")
	}
}

func TypeSwitch() {
	do(21)
	do("hello")
	do(true)
}
