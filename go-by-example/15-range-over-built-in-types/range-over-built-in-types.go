// range iterates over elements in a variety of built-in data structures. Let's see how to use range with some of the data structures we've already learned.

package main

import "fmt"

func main() {
	// here we use range to sum the numbers in a slice. Arrays work like this too.
	nums := []uint8{2, 3, 4}
	var sum uint8 = 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// range over arrays and slices provides both the index and value for each entry, Above we didn't need the index, so we ignored it with the blank identifier _. Sometimes we actually want the indexes though
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range over map iterates over key/value pairs.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// when you use range on a string, Go walks through each character not each byte
	// NOTE: c is a rune -> Go's word for a Unicode character
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
