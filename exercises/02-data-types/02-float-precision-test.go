package datatypes

import "fmt"

func FloatPrecisionTest() {
	a := 0.1
	b := 0.2

	/*
		Due to the IEEE 754 binary64 floating-point representation, decimal fractions like 0.1 and 0.2 cannot be represented exactly, leading to small rounding errors.Therefore, direct equality comparison (==) between floating point results and expected decimal value are unreliable. Use an epsilon-based comparison instead.
	*/

	fmt.Println(a + b)
}
