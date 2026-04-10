package variables

import "fmt"

func ModifyConstantError() {
	const age = 20
	// age = 30 - we can't do this because this is waht constant is for we can't reassign the constant value
	fmt.Println(age)
}
