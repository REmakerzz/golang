package main

import (
	"fmt"
	"strings"
)

func countWordOccurrences(text string) map[string]int {
	wordCount := make(map[string]int)

	words := strings.Fields(text)

	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func main() {
	text := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
	occurrences := countWordOccurrences(text)

	for word, count := range occurrences {
		fmt.Printf("%s: %d\n", word, count)
	}
}
