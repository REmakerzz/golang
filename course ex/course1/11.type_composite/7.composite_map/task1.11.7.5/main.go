package main

import (
	"fmt"
	"strings"
)

func filterSentence(sentence string, filter map[string]bool) string {
	map1 := make(map[string]bool)

	str := strings.Fields(sentence)
	var resultStr []string

	for _, word := range str {
		if map1[word] == filter[word] {
			resultStr = append(resultStr, word)
		}
	}
	return strings.Join(resultStr, " ")
}

func main() {
	sentence := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
	filter := map[string]bool{"ipsum": true, "elit": true}

	filtredSentence := filterSentence(sentence, filter)
	fmt.Println(filtredSentence)
}
