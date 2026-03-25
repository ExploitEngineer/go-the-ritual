/*
Named Return Values

Function return parameters can be named and treated as variables within function. Initialized to zero values. return statement without arguments returns current values of named parameters. Improves readability and enables easier refactoring but use judiciously.
*/

package main

import "fmt"

func wishBirthday(name string, age uint8) (wish string) {
	return fmt.Sprintf("Happry Birthday %s, You are now %d years old", name, age)
}

func NamedReturnValues() {
	wish := wishBirthday("whoami", 18)
	fmt.Println(wish)
}
