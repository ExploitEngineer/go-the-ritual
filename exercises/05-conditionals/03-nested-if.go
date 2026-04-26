package conditionals

import "fmt"

func NestedIf() {
	isLoggedIn := true
	isAdmin := false

	if isLoggedIn {
		if isAdmin {
			fmt.Println("Access Granted")
		} else {
			fmt.Println("Admin only")
		}
	} else {
		fmt.Println("Please login")
	}
}
