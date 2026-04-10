package datatypes

import "fmt"

func OverflowExperiment() {
	var num int8 = 127
	num++
	// int8 max is 127, incrementing causes overflow and wraps around to -128 due to two's complement representation.

	fmt.Println(num)
}
