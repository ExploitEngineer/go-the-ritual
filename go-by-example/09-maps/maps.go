// A map is Go's built in associative data type (hash table). It stores key-value pairs with fast average-time lookups.

package main

import "fmt"

func main() {
	// delcare a Map
	m := make(map[string]uint8) // a map with string keys and int values
	m["alice"] = 23
	m["bob"] = 30
	fmt.Println(m)

	// access values
	age := m["alice"]
	fmt.Println(age) // 23

	// if the key doesn't exists you get the zero
	age = m["charlie"]
	fmt.Println(age) // 0

	user := map[string]int{
		"rafay": 18,
		"bob":   0,
	}

	// ok is a boolean that is true if the key exists in the map, and false if it doesn't
	userAge, ok := user["charlie"]
	if !ok {
		fmt.Println("Key not found")
	} else {
		fmt.Println(userAge)
	}

	// how to iterate over a map
	for key, value := range m {
		fmt.Printf("%s: %d\n", key, value)
	}
}
