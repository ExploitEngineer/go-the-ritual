/*
break

Immediately exits innermost loop or switch statement. In nested loops, only exits immediate loop unless used with labels to break outer loops. Essential for early termination when conditions are met. Helps write efficient loops that don't continue unnecessarily.
*/

package main

import "fmt"

func Break() {
	for i := range 10 {
		if i == 5 {
			fmt.Println("Breaking out of loop")
			break // break here
		}
		fmt.Println("The value of i is", i)
	}
	fmt.Println("Exiting program")
}
