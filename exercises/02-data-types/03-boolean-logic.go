package datatypes

import "fmt"

func BooleanLogic() {
	isAdmin := true
	isLoggedIn := false

	fmt.Println(isAdmin && isLoggedIn) // false
	fmt.Println(isAdmin || isLoggedIn) // true
	fmt.Println(!isLoggedIn)           // true
}
