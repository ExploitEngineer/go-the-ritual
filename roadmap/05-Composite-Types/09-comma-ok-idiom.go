/*
Comma-Ok Idiom

Pattern for safely testing map key existence, type assertion success,
channel state, or error returns.

Syntax:
value, ok := something

ok is a boolean that tells whether the operation succeeded.
*/

package main

import "fmt"

func CommaOKIdiom() {
	// 1. Function return (error checking)

	n, err := divide(10, 2)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("result:", n)
	}

	// 2. Map key existence

	m := map[string]int{"a": 1}

	v, ok := m["a"]
	if ok {
		fmt.Println("value:", v)
	} else {
		fmt.Println("key not found")
	}

	// 3. Type assertion

	var i interface{} = "hello"

	s, ok := i.(string)
	if ok {
		fmt.Println("string:", s)
	} else {
		fmt.Println("not a string")
	}

	// 4. Channel closed check

	ch := make(chan int, 1)
	ch <- 10
	close(ch)

	val, ok := <-ch
	fmt.Println("value:", val, "open:", ok)

	val, ok = <-ch
	fmt.Println("value:", val, "open:", ok)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
