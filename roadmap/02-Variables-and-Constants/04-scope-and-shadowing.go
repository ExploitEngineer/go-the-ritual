/*
Scope determines variable accessibility from universe to block level. Variable shadowing occurs when a variable declared within certain scope (like a function or a block) has the same name as a variable in outer scope. The inner variable effectively *shadows* the outer variable, meaning that the inner variable is accessible in that scope while hte outer is temporarily hidden or inaccessible. Go has package, function, and block scopes.
Understanding prevents bugs from accidentally creating new variables.
*/

package main

import "fmt"

func ScopeAndShadows() {
	x := 10        // Outer variable
	fmt.Println(x) // Prints: 10

	if true {
		x := 5         // Inner variable (shadows the outer x)
		fmt.Println(x) // Prints: 5 (inner x)
	}

	fmt.Println(x) // Prints: 10 (outer x, accesible again)
}
