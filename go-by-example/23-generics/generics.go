package main

import "fmt"

type Stack[T any] struct {
	elements []T
}

func printSlice[T uint8 | string](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	myStack := Stack[int]{
		elements: []int{1, 2, 3},
	}

	newStack := Stack[string]{
		elements: []string{"golang", "typescript", "c"},
	}

	fmt.Println(myStack)
	fmt.Println(newStack)

	programmingLanguages := []string{"golang", "rust", "zig"}
	printSlice(programmingLanguages)

	nums := []uint8{1, 2, 3}
	printSlice(nums)
}
