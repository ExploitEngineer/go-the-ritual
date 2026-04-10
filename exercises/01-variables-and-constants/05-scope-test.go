package variables

import "fmt"

var name = "Global"

func ScopeTest() {
	name := "Local"
	fmt.Println(name)
}
