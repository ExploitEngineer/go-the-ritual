/*
if

Basic conditional statement for executing code based on boolean conditions. Supports optional initialization statement before condition check. No parentheses required around condition but braces mandatory. Can be chained with else if for multiple conditions. Foundation of control flow.
*/

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func If() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
