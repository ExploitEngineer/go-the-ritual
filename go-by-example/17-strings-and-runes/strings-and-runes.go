// Go strings are stored as read-only sequences of bytes like this: string -> [byte][byte byte][byte][byte]
// Text is encoded using UTF-8.
// A rune represents one Unicode character.

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Thai greeting "สวัสดี"
	// Looks like 5 characters,
	// but internally stored as multiple bytes.
	const s = "สวัสดี"

	// len() returns number of BYTES, not characters
	fmt.Println("Byte length:", len(s))

	// Print each byte (raw data)
	fmt.Println("\nBytes in the string:")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// Count actual characters (runes)
	fmt.Println("\nRune count:", utf8.RuneCountInString(s))

	// range automatically decodes UTF-8 into runes
	fmt.Println("\nIterating using range:")
	for index, r := range s {
		fmt.Printf("Character: %#U starts at byte %d\n", r, index)
	}

	// Manual rune decoding
	fmt.Println("\nManual UTF-8 decoding:")
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("Character: %#U starts at byte %d\n", r, i)
		checkRune(r)
		i += size
	}
}

// Rune literals use single quotes
func checkRune(r rune) {
	switch r {
	case 'ส':
		fmt.Println("→ Found Thai character 'ส'")
	case 'A':
		fmt.Println("→ Found letter A")
	}
}
