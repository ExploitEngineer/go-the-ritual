package datatypes

import "fmt"

func StringLengthTrap() {
	name := "Rafay"
	emoji := "🤍"

	fmt.Println(len(name))
	fmt.Println(len(emoji))

	// why may the emoji length surprise you ?
	// cause in go strings are immutable sequences of bytes and emoji can have many bytes and len function will return the len of the bytes in string not the rune count so we are getting the bytes count in string and that's why emoji has more bytes
}
