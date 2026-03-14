/*
for loop

Go's only looping construct, incredibly flexible for all iteration needs. Classic form: initialization, condition, post statements. Omit components for different behaviors (infinite, while-like). Use with break, continue, labels for nested loops. for range for convenient collection iteration.
*/

package main

import "fmt"

func ForLoop() {
	// simple for loop example
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break // loop terminated if i > 5
		}
		fmt.Printf("%d", i)
	}
	fmt.Printf("\nloop ended")

	// continue keyword example
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("Odd: %d", i)
	}
}
