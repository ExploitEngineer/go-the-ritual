package conditionals

import "fmt"

func NumberSignDetector() {
	num := 1
	if num == 0 {
		fmt.Println("zero")
	} else if num < 0 {
		fmt.Println("negative")
	} else {
		fmt.Println("positive")
	}
}
