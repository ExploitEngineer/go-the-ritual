package main

import "fmt"

// The (uint8, uint8) in this function signature shows that the function returns 2 ints.
func vals() (uint8, uint8) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
