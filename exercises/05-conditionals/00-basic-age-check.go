// Package conditionals include all the exercises related to conditionals
package conditionals

import "fmt"

func BasicAgeCheck() {
	age := 18

	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}
}
