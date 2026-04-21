package structs

import "fmt"

type person struct {
	Name string
	Age  uint8
}

func RealMiniSystem() {
	users := map[string]person{}

	username := "rafay"

	if u, ok := users[username]; ok {
		u.Age++
		users[username] = u
	} else {
		users[username] = person{
			Name: "Abdul Rafay",
			Age:  18,
		}
	}

	fmt.Println(users)
}
