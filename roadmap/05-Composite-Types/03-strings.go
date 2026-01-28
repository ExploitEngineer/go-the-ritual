/*
Immutable sequences of bytes representing UTF-8 encoded text. String operations create new strings rather than modifying existing ones. Iterate by bytes (indexing) or runes (for range). Convert between strings and byte slices. Understanding strings helps with text manipulation and performance.
*/

package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func Strings() {
	/*
		String Literals
		String literals are enclosed in double quotes, and they allow you to use escape sequences like \n for newline and \t for tab.
	*/
	message := "Hello, Go!"
	fmt.Println(message)

	/*
		Raw String Literals
		Raw string literals are enclosed in backticks and are useful when you want to include special characters or multiline strings without worrying about escape sequences.
	*/
	rawMessage := `This is a raw string literal.
	It can span multiple lines
	without escape sequences like \n.`
	fmt.Println(rawMessage)

	/*
		String Concatenation
		String concatenation in Go is simple. You can use the + operator to combine strings.
	*/
	firstName := "John"
	lastName := "Doe"
	fullName := firstName + " " + lastName
	fmt.Println(fullName)

	/*
		String Conversions
		To convert other data types to strings, you can use the fmt.Sprintf function or type conversion.
	*/
	age := 30
	ageStr := fmt.Sprintf("Age: %d", age)
	fmt.Println(ageStr)

	score := 95.5
	scoreStr := strconv.FormatFloat(score, 'f', 2, 64)
	fmt.Println("Score: " + scoreStr)

	/*
		Calculating String Length
		In Go, the built-in len function is used to calculate the length of a string, but it doesn't work as expected when dealing with multi-byte characters
	*/
	text := "Hello, 你好"
	length := len(text)                        // Incorrect length
	runeLength := utf8.RuneCountInString(text) // Correct length
	fmt.Printf("Incorrect Length: %d\n", length)
	fmt.Printf("Correct Length: %d\n", runeLength)
}
