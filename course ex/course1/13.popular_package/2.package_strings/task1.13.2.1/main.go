package main

import (
	"fmt"
	"strings"
)

func CountWordsInText(txt string, words []string) map[string]int {
	if len(txt) == 0 {
		return map[string]int{}
	}
	if len(words) == 0 {
		return map[string]int{}
	}
	lowerCaseTxt := strings.Fields(strings.ToLower(txt))
	wordCount := make(map[string]int)

	for _, word := range lowerCaseTxt {
		wordCount[word]++

	}

	result := make(map[string]int)

	for _, word := range words {
		wordLower := strings.ToLower(word)
		result[wordLower] = wordCount[wordLower]
	}

	return result
}

func main() {
	txt := `Lorem ipsum dolor sit amet , consectetur adipiscing elit. Donec a diam lectus. Sed sit amet ipsum mauris. 
        Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit. Donec et mollis dolor. 
	Praesent et diam eget libero egestas mattis sit amet vitae augue.`
	words := []string{"sit", "amet", "lorem"}

	result := CountWordsInText(txt, words)

	fmt.Println(result)

}
