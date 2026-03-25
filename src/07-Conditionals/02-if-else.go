/*
if-else

Basic conditional statements for binary decision making. if tests condition, else handles alternative path. Can include optional initialization statement. No parentheses needed around condition but braces required. Foundation of program control flow.
*/

package main

import (
	"fmt"
	"math"
)

func doPow(x, n, lim float64) float64 {
	// v is initialized here and is visible in all blocks below
	if v := math.Pow(x, n); v < lim {
		return v
	} else if v == lim {
		fmt.Printf("Exactly at the limit: %g\n", v)
		return v
	} else {
		fmt.Printf("%g is over the limit %g\n", v, lim)
	}

	// v is no longer accessible here; the scope has ended
	return lim
}

func IfElse() {
	fmt.Println("Result 1:", doPow(3, 2, 10)) // 9 < 10
	fmt.Println("Result 2:", doPow(2, 3, 8))  // 8 == 8
	fmt.Println("Result 3:", doPow(3, 3, 10)) // 27 > 10
}
