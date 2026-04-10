package variables

import "fmt"

func BuildUserCard() {
	name := "rafay"
	age := 20
	country := "Pakistan"

	fmt.Printf("User Info:\n")
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Country: %s\n", country)
}
