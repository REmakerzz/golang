package main

import (
	"fmt"
	"strings"
	"unicode"
)

func CountVowels(str string) int {
	vowels := "аеёиоуыэюяАЕЁИОУЫЭЮЯaeiouyAEIOUY"

	var count int

	for _, char := range str {
		if unicode.IsLetter(char) && strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}

func main() {
	count := CountVowels("Привет, мир!")
	fmt.Println(count)

	count = CountVowels("Hello, world!")
	fmt.Println(count)

	count = CountVowels("12345")
	fmt.Println(count)

	count = CountVowels("")
	fmt.Println(count)

	count = CountVowels("A")
	fmt.Println(count)
}
