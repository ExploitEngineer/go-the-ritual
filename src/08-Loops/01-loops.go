/*
Loops

Go has only one looping construct: the flexible for loop. Basic form has initialization, condition, post statement. Supports for range for arrays, slices, maps, strings, channels. Can create infinite loops or while-style loops. Control with break and continue.
*/

package main

import "fmt"

func Loops() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
