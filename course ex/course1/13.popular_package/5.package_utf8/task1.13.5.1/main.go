package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	fmt.Println(countUniqueUTF8Chars("Hello, 世界!"))
}

func countUniqueUTF8Chars(s string) int {
	uniqueChars := make(map[rune]struct{})

	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		uniqueChars[r] = struct{}{}
		s = s[size:]
	}

	return len(uniqueChars)
}
