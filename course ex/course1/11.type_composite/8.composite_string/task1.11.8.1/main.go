package main

import (
	"fmt"
	"unicode/utf8"
)

func countBytes(s string) int {
	return len(s)
}

func countSymbols(s string) int {
	return utf8.RuneCountInString(s)
}

func main() {
	var message string = "Привет, мир!"

	bytes := countBytes(message)
	fmt.Println(bytes)
	symbols := countSymbols(message)
	fmt.Println(symbols)
}
