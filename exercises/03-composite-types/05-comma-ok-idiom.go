package compositetypes

import "fmt"

func CommaOkIdiom() {
	m := map[string]int{"a": 1}

	// what are the value and ok?
	// the value is the value to key we passed b and ok is the bool it returns false if that key doesn't exist in map and return true if it does.
	value, ok := m["b"]

	if !ok {
		fmt.Println("value doesn't exist in map")
	} else {
		fmt.Println(value)
	}
}
