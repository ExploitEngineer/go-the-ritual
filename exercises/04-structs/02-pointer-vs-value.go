package structs

func receiveByValue(u User) {
	u.Name = "John"
}

func receiveByPointer(u *User) {
	u.Age = 17
}

func PointerVsValues() {
	user := User{Name: "Rafay", Age: 20}

	// why does one change original and other not?
	// in the funtion where we are passing by value is not chaning the underlying struct cause it is only changing the one which we created.
	// but the one which we are passing by the pointer is passing the value as the reference the underlying actual struct memory address so that is why it is changing the original value of it.
	receiveByValue(user)
	receiveByPointer(&user)
}
