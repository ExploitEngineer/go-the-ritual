/*
switch

Clean way to compare variable against multiple values and execute corresponding code blocks. No break statements needed (no fall-through by default). Works with any comparable type, supports multiple values per case, expression/type switches. More readable than if-else chains.
*/

package main

import (
	"fmt"
	"time"
)

func Switch() {
	i := 2
	fmt.Println("Write ", i, " as ")

	// Here's a basic switch
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// You can use commas to separate multiple expressions in the same case statement. We use the optional default case in this example as well.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// switch without an expression is an alternate way to express if/else logic. Here we also show how the case expression can be non-constants.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// A type switch compares types instead of values. You can use this to discover the type of an interface value. In this example, the variable t will have the type corresponding to its clause.
	whatAmI := func(i any) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm am bool")
		case int:
			fmt.Println("I'm am int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(10)
	whatAmI("hey")
}
