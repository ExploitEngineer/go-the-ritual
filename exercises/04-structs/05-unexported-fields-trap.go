package structs

import (
	"encoding/json"
	"fmt"
)

type user struct {
	name string
	age  uint8
}

func UnexpectedFieldsTrap() {
	u := user{name: "rafay", age: 20}
	s, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
		return
	}

	// what does it fail/return empty JSON ?
	// Fields starting lowercase are private to package.
	fmt.Println(s) // {}
}
