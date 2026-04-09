package main

import "fmt"

/*
Introduction to Runes

A rune in Go represents a single Unicode code point.

Technically:
	type rune = int32

This means a rune is just a number underneath whose purpose is
to represent one Unicode character.

Examples:
	'A' = 65
	'a' = 97
	'¥' = 165
	'你' = 20320

Go uses runes because ASCII was too limited and could not support
many languages/symbols. Unicode solved this by assigning every
character a unique numeric value, and Go stores those values as runes.
*/

/*
Basic Usage of Runes

Runes are written using single quotes.

They can be initialized directly, with unicode escape sequences,
or by converting numbers into runes.
*/
func BasicRunes() {
	letter := 'A'
	yen := '¥'
	unicodeA := '\u0041'
	numberRune := rune(66)

	fmt.Println(letter) // 65
	fmt.Println(yen)    // 165

	fmt.Printf("%c\n", letter)     // A
	fmt.Printf("%c\n", yen)        // ¥
	fmt.Printf("%c\n", unicodeA)   // A
	fmt.Printf("%c\n", numberRune) // B

	// Since rune is numeric, arithmetic works
	fmt.Printf("%c\n", 'A'+1) // B
}

/*
Text Analyzer Use Case

One common use of runes is analyzing text character-by-character.

Below is a simple analyzer that counts:
  - total characters
  - uppercase letters
  - lowercase letters
  - words
  - lines

We use []rune here because runes let us inspect individual characters directly.
*/
func TextAnalyzer() {
	text := []rune{
		'H', 'e', 'l', 'l', 'o', ' ',
		'W', 'O', 'R', 'L', 'D',
		'\n',
		'G', 'o',
	}

	var (
		characters int
		uppercase  int
		lowercase  int
		words      int
		lines      int = 1
	)

	inWord := false

	for _, r := range text {
		characters++

		if r >= 'A' && r <= 'Z' {
			uppercase++
		}

		if r >= 'a' && r <= 'z' {
			lowercase++
		}

		if r == '\n' {
			lines++
		}

		if r != ' ' && r != '\n' && !inWord {
			words++
			inWord = true
		}

		if r == ' ' || r == '\n' {
			inWord = false
		}
	}

	fmt.Println("Characters:", characters)
	fmt.Println("Uppercase :", uppercase)
	fmt.Println("Lowercase :", lowercase)
	fmt.Println("Words     :", words)
	fmt.Println("Lines     :", lines)
}
