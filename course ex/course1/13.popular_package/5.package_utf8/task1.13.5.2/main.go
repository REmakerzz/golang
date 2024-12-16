package main

import (
	"fmt"
	"unicode"
)

func countRussianLetters(s string) map[rune]int {
	counts := make(map[rune]int)

	for _, char := range s {
		if unicode.Is(unicode.Cyrillic, char) {
			counts[unicode.ToLower(char)]++
		}
	}
	return counts
}

func main() {
	s := countRussianLetters("Привет, мир!")
	for key, value := range s {
		fmt.Printf("%c: %d \n", key, value)
	}
}
