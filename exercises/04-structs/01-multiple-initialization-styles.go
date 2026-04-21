package structs

import "fmt"

func MultipleInitializationStyles() {
	// positional
	u := User{"rafay", 18}

	// named
	user := User{Name: "Rafay", Age: 18}
	fmt.Println(u)
	fmt.Println(user)
}
