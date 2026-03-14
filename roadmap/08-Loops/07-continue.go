/*
continue

Skips rest of current iteration and jumps to next loop iteration. Only affects innermost loop unless used with labels. Useful for filtering elements, handling special cases early, avoiding nested conditionals. Makes loops cleaner and more efficient.
*/

package main

import "fmt"

func Continue() {
	for i := range 10 {
		if i == 5 {
			fmt.Println("Continuing loop")
			continue // break here
		}
		fmt.Println("The value of i is", i)
	}
	fmt.Println("Exiting program")
}
