// Package structs include all the exercises related to structs
package structs

import "fmt"

type User struct {
	Name string
	Age  int
}

func BasicStructCreation() {
	user := User{Name: "Abdul Rafay", Age: 18}
	fmt.Println("Name: %s, Age: %d", user.Name, user.Age)
}
