/*
Go provides two main ways to declare variables: using `var` and using the short declaration operator `:=`.

The `var` keyword is used for explicit variable declarations. You can use it to define a variable with or without assigning a value. If no value is provided, Go assigns a default zero value based on the variable type.
`var` can be used both inside and outside functions.

The `:=` syntax is a shorthand for declaring and initializing a variable. It infers the type from the value and can only be used inside functions. This is a quick and convevient way to create variables without explicitly mentioning their types.
*/

package main

import "fmt"

func VarVsShort() {
	var i, j uint8 = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
