package compositetypes

import "fmt"

func MissingKeyPitfall() {
	m := map[string]int{}

	// What is output and why can this be dangerous?
	// The value is 0
	fmt.Println(m["missing"])
}
