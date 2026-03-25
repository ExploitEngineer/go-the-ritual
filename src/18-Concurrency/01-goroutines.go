/*
Goroutines

Lightweight threads managed by Go runtime enabling concurrent function execution. Created with go keyword prefix. Minimal memory overhead, can run thousands/millions concurrently. Runtime handles scheduling across CPU cores. Communicate through channels, fundamental to Go's concurrency.
*/
package main

import (
	"fmt"
	"time"
)

/*
WHAT IS A GOROUTINE?

JS/TS → Go equivalents:
  setTimeout(() => {}, 0)   → go someFunc()
  Promise / async-await     → goroutine + channels (next topic)
  worker threads            → goroutines (but way cheaper)

A goroutine is a function that runs concurrently alongside other functions.
Think of it like saying "start this function, but don't wait for it — keep going."

The keyword `go` before any function call is all you need.
*/

/*
WHY ARE GOROUTINES BETTER THAN THREADS?

Threads (JS worker threads, Java threads, etc.):
  - each thread = ~1MB of memory
  - OS manages them (slow to create/destroy)
  - fixed stack size

Goroutines:
  - each goroutine = ~2KB of memory (500x cheaper than threads)
  - Go runtime manages them (fast)
  - stack grows and shrinks automatically as needed
  - you can run thousands (even millions) at the same time
*/

/*
THE MOST IMPORTANT RULE ABOUT GOROUTINES:

When you call `go someFunc()`:
  1. it starts immediately
  2. control does NOT wait for it to finish
  3. the next line runs right away

And this is critical:
  If main() finishes, ALL goroutines are killed — even if they are still running.

This is why the program below does NOT print "Hello world goroutine":
*/

func hello() {
	fmt.Println("Hello world goroutine")
}

func goroutineBasic() {
	go hello()
	fmt.Println("main function")
	/*
	   What happens here step by step:
	     1. go hello()  → goroutine starts BUT we don't wait for it
	     2. fmt.Println → prints "main function" immediately
	     3. function ends → goroutine never got a chance to run

	   JS comparison:
	     setTimeout(() => console.log("hello"), 0)
	     console.log("main")
	     → "main" prints first, same idea
	*/
}

/*
THE FIX (TEMPORARY HACK — NOT FOR REAL CODE):

We can use time.Sleep to pause main long enough for the goroutine to finish.
This is just to understand how goroutines work.
In real code you would use channels (next topic) to do this properly.
*/

func goroutineWithSleep() {
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("main function")
	/*
	   Now what happens:
	     1. go hello()         → goroutine starts
	     2. time.Sleep(1s)     → main pauses for 1 second
	     3. hello() runs       → prints "Hello world goroutine"
	     4. sleep ends         → prints "main function"
	*/
}

/*
MULTIPLE GOROUTINES RUNNING AT THE SAME TIME:

You can start as many goroutines as you want.
They all run concurrently — Go's runtime decides which one runs when.
*/

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func multipleGoroutines() {
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
	/*
	   Both goroutines run at the same time.
	   numbers() prints every 250ms: 1, 2, 3, 4, 5
	   alphabets() prints every 400ms: a, b, c, d, e
	   They interleave because they run concurrently:

	   output: 1 a 2 3 b 4 c 5 d e main terminated

	   Timeline:
	   0ms    → both goroutines start
	   250ms  → numbers prints 1
	   400ms  → alphabets prints a
	   500ms  → numbers prints 2
	   750ms  → numbers prints 3
	   800ms  → alphabets prints b
	   1000ms → numbers prints 4
	   1200ms → alphabets prints c
	   1250ms → numbers prints 5
	   1600ms → alphabets prints d
	   2000ms → alphabets prints e
	   3000ms → main prints "main terminated"
	*/
}

func Goroutines() {
	fmt.Println("--- basic (goroutine gets killed before running) ---")
	goroutineBasic()

	fmt.Println("\n--- with sleep (goroutine gets time to run) ---")
	goroutineWithSleep()

	fmt.Println("\n--- multiple goroutines running concurrently ---")
	multipleGoroutines()
}
