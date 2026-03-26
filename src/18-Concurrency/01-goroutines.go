/*
Goroutines are a lightweight form of concurrent programming in Golang. They are used to handle multiple tasks simultaneously within a single program. Goroutines are created using the ‘go’ keyword, followed by a function call. When a Goroutine is created, it runs in parallel with the main function, and both the Goroutine and the main function can run independently, executing their respective code.
*/

package main

import (
	"fmt"
	"time"
)

func printMessage(message string) {
	for i := range 10 {
		fmt.Printf("%d: %s\n", i, message)
		time.Sleep(1 * time.Second)
	}
}

func Goroutines() {
	go printMessage("Hello from Goroutine")
	fmt.Println("Main function")
	time.Sleep(11 * time.Second)
}
