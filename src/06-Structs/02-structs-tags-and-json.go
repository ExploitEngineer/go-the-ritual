/*
Struct Tags & JSON

Struct tags provide metadata about struct fields.
They are written using backticks (`) and follow a key:"value" format.

JSON tags are used by the encoding/json package to control how
struct fields appear in JSON.

They allow you to:
- Rename fields
- Omit empty fields
- Ignore fields completely

Example tag:
json:"name,omitempty"
*/

package main

import (
	"encoding/json"
	"fmt"
)

// NOTE: Defining a Struct with JSON Tags

type Person struct {
	Username    string `json:"username"`
	Email       string `json:"email,omitempty"`
	SignInCount uint   `json:"sign_in_count"`
	IsActive    bool   `json:"is_active"`
	Password    string `json:"-"` // ignored in JSON
}

func StructTagsJSON() {
	// NOTE: Creating a Struct Instance

	user := Person{
		Username:    "alice",
		Email:       "alice@example.com",
		SignInCount: 5,
		IsActive:    true,
		Password:    "supersecret",
	}

	fmt.Println("Original Struct:")
	fmt.Println(user)

	/*
		NOTE: json.Marshal()

		Converts a Go struct into JSON format.
		Returns JSON as []byte and an error.
	*/

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nStruct converted to JSON:")
	fmt.Println(string(jsonData))

	/*
		NOTE: json.MarshalIndent()

		Works like Marshal but formats the JSON
		with indentation for readability.
	*/

	prettyJSON, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nPretty Printed JSON:")
	fmt.Println(string(prettyJSON))

	/*
		NOTE: json.Unmarshal()

		Converts JSON data back into a Go struct.
		The second argument must be a pointer
		so the function can modify the struct.
	*/

	jsonInput := `{
		"username": "bob",
		"email": "bob@example.com",
		"sign_in_count": 3,
		"is_active": true
	}`

	var newUser User

	err = json.Unmarshal([]byte(jsonInput), &newUser)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nJSON converted back to Struct:")
	fmt.Println(newUser)
}
