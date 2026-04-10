package datatypes

import "fmt"

func ComplexNumber() {
	c := complex(2, 3)

	fmt.Println(c)
	fmt.Println(real(c))
	fmt.Println(imag(c))
}
