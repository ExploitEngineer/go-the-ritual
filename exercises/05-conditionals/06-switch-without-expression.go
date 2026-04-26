package conditionals

import "fmt"

func SwitchWithoutExpression() {
	age := 10
	switch {
	case age < 13:
		fmt.Println("child")
	case age < 18:
		fmt.Println("minor")
	case age >= 18:
		fmt.Println("Adult")
	}
}
