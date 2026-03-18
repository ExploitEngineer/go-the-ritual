/*
Call by Value

Go creates copies of values when passing to functions, not references to originals. Applies to all types including structs and arrays. Provides safety but can be expensive for large data. Use pointers, slices, maps for references. Critical for performance optimization.
*/

package main

import "fmt"

type User struct {
	Name string
	Age  uint8
}

// This function receives a COPY of the struct
func udpateAgeByValue(u User) {
	u.Age = u.Age + 1
	fmt.Printf("Inside Function (Copy): %d\n", u.Age)
}

func CallByValue() {
	user1 := User{Name: "Alice", Age: 25}

	fmt.Printf("Before Function: %d\n", user1.Age)

	udpateAgeByValue(user1)

	fmt.Printf("After Function: %d\n", user1.Age)
	// Output is still 25 because the function only touched a copy!
}
