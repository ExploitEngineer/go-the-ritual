// Go supports anonymous functions, which can form closures. Anonymous functions are useful when you want to define a function inline without having to name it.

package main

import "fmt"

// This function intSeq returns another function, which we define anonymously in the body of intSeq. The returned function closes over the variable 1 to form a closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	newInts := intSeq()
	fmt.Println(newInts()) // 1
}
