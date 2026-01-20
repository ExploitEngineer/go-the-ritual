/*
In Go, strings are read-only slice of bytes representing characters. Let's start by understanding the different ways to declare and initialize strings using string literals and raw string literals.

- String Literal
String literals are enclosed in double quotes, and they allow you to use escape sequences like \n for newline or \t for tab.

- Raw String Literals
Raw string literals are enclosed in backticks and are useful when you want to include special characters or multiline strings without worrying about escape sequences.

- String Concatenation
String concatenation in Go is simple. You can use the + operator to combine strings.

- String Conversions
To convert other data types to strings, you can use the fmt.Sprintf function or type conversion.
*/

package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func String() {
	// here message is a string literal
	message := "Hello, Go!"
	fmt.Println(message)

	// Raw String Literal
	rawMessage := `This is a raw string literal.
	It can spcan multiple lines
	without escape sequences like \n.`
	fmt.Println(rawMessage)

	// String Concatenation
	firstName := "Abdul"
	lastName := "Rafay"
	fullName := firstName + " " + lastName
	fmt.Println(fullName)

	// String Conversion
	age := 18
	ageStr := fmt.Sprintf("Age: %d", age)
	fmt.Println(ageStr)

	score := 95.5
	scoreStr := strconv.FormatFloat(score, 'f', 2, 64)
	fmt.Println("Score: " + scoreStr)

	// Calculating String Length
	// In Go, the built-in len function is used to calculate the length of the string, but it doesn't work as expected when dealing with multi-byte characters.
	text := "Hello, 你好"
	length := len(text)                        // Incorrect length
	runeLength := utf8.RuneCountInString(text) // Correct length
	fmt.Printf("Incorrect Length: %d\n", length)
	fmt.Printf("Correct Length: %d\n", runeLength)
}
