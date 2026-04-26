package conditionals

import "fmt"

func LoginValidation() {
	username := "rafay"
	password := "12345"

	if username == "rafay" && password == "12345" {
		fmt.Println("Login successful")
	} else {
		fmt.Println("Invaild credentails")
	}
}
