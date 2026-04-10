package variables

import "fmt"

func ConstantsPractice() {
	const (
		Pi      float64 = 3.14
		AppName string  = "anything"
		Version float64 = 1.2
	)

	fmt.Println(Pi, AppName, Version)
}
