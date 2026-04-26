package conditionals

import "fmt"

func MiniGradeSystem() {
	marks := 76

	if marks > 90 {
		fmt.Println("A")
	} else if marks > 80 {
		fmt.Println("B")
	} else if marks > 70 {
		fmt.Println("C")
	} else if marks > 60 {
		fmt.Println("D")
	} else {
		fmt.Println("Fail")
	}
}
