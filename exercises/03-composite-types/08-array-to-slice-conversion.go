package compositetypes

import "fmt"

func ArrayToSlice() {
	arr := [3]int{1, 2, 3}

	s := arr[:]
	s1 := []int{arr[0], arr[1], arr[2]}

	fmt.Println(s)
	fmt.Println(s1)
}
