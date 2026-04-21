package structs

import "fmt"

type Address struct {
	City    string
	Country string
}

type Person struct {
	Name    string
	Address Address
}

func NestedStructs() {
	p := Person{
		Name: "rafay",
		Address: Address{
			City:    "isb",
			Country: "pakistan",
		},
	}

	fmt.Println(p.Name)
	fmt.Println(p.Address.City)
}
