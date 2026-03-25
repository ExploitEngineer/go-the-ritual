/*
Structs
Custom data types grouping related fields under single name. Similar to classes but methods defined separately. Create complex data models, organize information, define application data structure. Access fields with dot notation, pass to functions. Fundamental for object-oriented designs.
*/

package main

import "fmt"

/*
NOTE: Defining a Struct
Structs are defined using the type and struct keywords. Here's an example of a simple struct definition:
*/

type User struct {
	Username    string
	Email       string
	SignInCount uint
	IsActive    bool
}

func Structs() {
	/*
		NOTE:Initializing a Struct
		Structs can be initialized in various ways.
	*/

	// Initializing with Field Names
	user1 := User{
		Username:    "alice",
		Email:       "alice@example.com",
		SignInCount: 1,
		IsActive:    true,
	}

	fmt.Println(user1)

	/*
		Initializing with Default Values
		If some fields are not specified, they are initialized to their zero values for the respective types.
	*/
	user2 := User{
		Username: "bob",
	}

	fmt.Println(user2)

	/*
		Initializing with a Pointer
		A struct can also be initialized using a pointer.
	*/
	user3 := &User{
		Username: "charlie",
		Email:    "charlie@example.com",
	}

	fmt.Println(user3)
}
