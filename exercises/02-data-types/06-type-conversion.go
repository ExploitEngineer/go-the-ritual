package datatypes

import "fmt"

func TypeConversion() {
	var age int = 18
	// converting float to int truncates decimal part, does not round
	floatAge := float64(age)
	fmt.Println(floatAge)

	var price float64 = 99.99
	intPrice := int(price)
	fmt.Println(intPrice)
}
