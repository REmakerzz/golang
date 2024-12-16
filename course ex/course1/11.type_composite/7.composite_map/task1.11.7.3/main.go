package main

import (
	"fmt"
	"strings"
)

func createUniqueText(text string) string {
	map1 := make(map[string]bool)

	var resultStr []string

	str := strings.Fields(text)
	for _, word := range str {
		if !map1[word] {
			map1[word] = true
			resultStr = append(resultStr, word)
		}
	}
	return strings.Join(resultStr, " ")
}

func main() {
	text := "bar bar bar foo foo baz"
	fmt.Println(createUniqueText(text))
}
