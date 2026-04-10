package variables

import "fmt"

func BlockScope() {
	if true {
		x := 0
		fmt.Println(x)
	}

	// fmt.Println(x) - x doesn't exist here it is only inside the block scope
}
