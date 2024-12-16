package main

import "fmt"

func ReplaceSymbols(str string, old rune, newChar rune) string {
	sliceRuneStr := []rune(str)

	for i, char := range sliceRuneStr {
		if char == old {
			sliceRuneStr[i] = newChar
		}
	}
	return string(sliceRuneStr)
}

func main() {
	result := ReplaceSymbols("Hello, world!", 'o', '0')
	fmt.Println(result)
}
