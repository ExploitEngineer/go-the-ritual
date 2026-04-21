package structs

import (
	"encoding/json"
	"fmt"
)

type JSONUser struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func JsonTags() {
	u := JSONUser{Name: "rafay", Age: 20}
	data, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))

	// without tags
	// {"Name":"rafay","Age":20}

	// with tags
	// {"name":"rafay","age":20}
}
