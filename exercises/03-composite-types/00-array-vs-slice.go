// Package compositetypes have all exercises related to compositetypes in go
package compositetypes

import "fmt"

func ArrayVsSlice() {
	arr := [5]int{1, 2, 3, 4, 5}

	// a slice is a struct containing pointer + len + cap
	s := []int{1, 2, 3, 4, 5}

	arr[0] = 10
	s[0] = 20

	fmt.Println(arr[0])
	fmt.Println(s[0])

	fmt.Println(arr)
	fmt.Println(s)
}
