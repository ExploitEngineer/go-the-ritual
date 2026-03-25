/*
Iterating Maps

Use for range to iterate over maps, returns key and value pairs. Iteration order is random for security reasons. Use blank identifier _ to ignore key or value. Cannot modify map during iteration unless creating new map. Safe to delete during iteration.
*/

package main

import "fmt"

func IteratingMaps() {
	sampleMap := map[string]int{
		"apple":  2,
		"banana": 5,
		"cherry": 7,
	}

	for key, value := range sampleMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
