/*
Built-in associative data type mapping keys to values. References types created with `make(map[KeyType]ValueType)` or map literals. Keys must be comparable types. Support insertion, deletion, lookup operations. Check existence with comma ok idiom: `value, ok := map[key]`
*/

package main

import "fmt"

func Maps() {
	// declare a map
	m := make(map[string]int) // a map with string keys and int values
	m["alice"] = 23
	m["bob"] = 30
	fmt.Println(m) // map[alice:23 bob:30]

	// access map values
	age := m["alice"]
	fmt.Println(age) // 23

	/*
		iterate over a map
		Iterating over a map means going through all key-value pairs in the map, one at a time. You do this with a foor loop and the range keyword:
	*/
	for key, value := range m {
		fmt.Printf("%s: %d\n", key, value)
	}
}
