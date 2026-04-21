package structs

import "fmt"

type str struct {
	name string
}

func StructComparison() {
	a := str{name: "rafay"}
	b := str{name: "rafay"}

	// when does this work and when does it fail?
	// it works when we use the same struct with the same vaules and compare both of them but if we compare two separate structs with same vaules it will throw an error
	fmt.Println(a == b) // true
}
