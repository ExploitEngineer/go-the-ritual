package compositetypes

import "fmt"

func SliceReferenceTrap() {
	a := []int{1, 2, 3}
	b := a

	b[0] = 100

	// why did a also change ?
	fmt.Println(a)
	fmt.Println(b)

	// a and b are two slice headers pointing to the SAME underlying array
}
