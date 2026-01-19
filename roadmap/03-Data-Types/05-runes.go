package main

/*
ASCII, Unicode, UTF & Runes

1. ASCII (The Old World)

ASCII = American Standard Code for Information Interchange

- Created in the early days of computing
- Uses **7 bits** (0–127)
- Can represent only:
  - English letters (A–Z, a–z)
  - Digits (0–9)
  - Basic symbols (!, @, #, etc.)
  - Control characters (\n, \t, etc.)

Example:
  'A' = 65
  'a' = 97
  '0' = 48

Problem with ASCII:
- It CANNOT represent:
  - Arabic
  - Chinese
  - Japanese
  - Emojis

ASCII is **too small** for the real world.

2. Unicode (The Global Solution)

Unicode is NOT an encoding.
Unicode is a **global character dictionary**.

Unicode:
- Assigns a unique NUMBER to every character
- This number is called a **code point**

Examples:
  'A'  -> U+0041
  '¥'  -> U+00A5
  '💞' -> U+1F49E

Unicode is a **superset of ASCII**:
- ASCII characters keep the same values

Important:
Unicode only defines numbers.
It does NOT define how those numbers are stored in memory.

3. UTF Encodings (Storage Rules)

UTF = Unicode Transformation Format

UTF answers the question:
"How do we store Unicode code points in memory?"

Common UTFs:
- UTF-8   (Go uses this)
- UTF-16
- UTF-32

UTF-8 (Very Important for Go)

UTF-8 properties:
- Variable length encoding
- Uses **1 to 4 bytes per character**

Rules:
- ASCII characters → 1 byte
- Other characters → 2 to 4 bytes

Examples:
  'A'   -> 1 byte
  '¥'   -> 2 bytes
  '💞'  -> 4 bytes

Why UTF-8 won?
- Backward compatible with ASCII
- Memory efficient
- Dominates the internet

Go strings are:
- Immutable
- UTF-8 encoded byte sequences

4. Rune (Go’s Character Type)

In Go:

  byte = alias for uint8  (raw byte)
  rune = alias for int32  (Unicode code point)

A rune represents:
- ONE Unicode character
- Stored as its Unicode code point value

Example:
  'A'  -> rune(65)
  '¥'  -> rune(165)
  '💞' -> rune(128158)

Rune exists because:
- Strings are bytes
- Humans think in characters

5. Declaring Runes
*/

func Runes() {
	// Rune literal (recommended)
	yen := '¥'

	// Integer (represents Unicode code point)
	japYen := 165

	// Hexadecimal Unicode value
	hexYen := 0x00A5

	// Unicode escape
	unicodeYen := '\u00A5'

	// Octal escape
	octalYen := '\245'

	_ = yen
	_ = japYen
	_ = hexYen
	_ = unicodeYen
	_ = octalYen
}

/*
6. Printing Runes

- %c  → prints rune as character
- string(rune) → converts rune to string

7. Strings vs Runes (CRITICAL)

Strings in Go are:
- UTF-8 encoded
- Sequence of BYTES

Indexing a string gives a BYTE, NOT a character.

Example:
  str := "💞 code"

  str[0] -> first byte (NOT 💞)

This is why slicing strings is dangerous
for Unicode text.

8. Correct Way: range

Using `range` on a string:
- Decodes UTF-8
- Returns runes (characters)

9. Mental Model (LOCK THIS)

byte  → storage unit
rune  → character
string → bytes (UTF-8)
range  → character-safe iteration

10. Why Runes Matter

Runes are essential for:
- Multilingual text
- Emojis
- Correct string processing
- Security-sensitive parsing

If you ignore runes, your program WILL break
for non-English input.
*/
