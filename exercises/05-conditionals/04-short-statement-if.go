package conditionals

import "fmt"

func ShortStatementIf() {
	// in golang we can also initialize variable in if statement
	if score := 85; score >= 80 {
		fmt.Println("Passed with destinction")
	}
}
