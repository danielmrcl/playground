package examples

import (
	"fmt"
	"unicode/utf8"
)

// Strings and Runes examples
func StringsExamples() {
	s := "Olá Mundo!"
	runeVal, width := utf8.DecodeRuneInString(s[2:]) // rune of 'á' char
	fmt.Println("string char:", s[2:2+width]) // index 2 to 4. (two bytes to rune)
	fmt.Printf("rune value: %#U\n", runeVal)
	fmt.Println("rune width:", width)

	if runeVal == 'á' {
		fmt.Println("found á!")
	}
}
