/*
Pointer Basics

Variables storing memory addresses of other variables. Declared with *Type, dereferenced with *ptr, address obtained with &var. Enable efficient memory usage and allow functions to modify caller's data. Essential for performance and reference semantics.
*/

package main

import "fmt"

func PointersBasics() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // dividie j through the pointer
	fmt.Println(j) // see the new value of j
}
