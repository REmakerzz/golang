package main

import (
	"fmt"
	"strings"
)

func ConcatenateStrings(sep string, str ...string) string {
	var evenStr []string
	var oddStr []string

	for i, s := range str {
		if i%2 == 0 {
			evenStr = append(evenStr, s)
		} else {
			oddStr = append(oddStr, s)
		}
	}

	evenResult := strings.Join(evenStr, sep)
	oddResult := strings.Join(oddStr, sep)

	evenResult = strings.TrimSuffix(evenResult, sep)
	oddResult = strings.TrimSuffix(oddResult, sep)
	return fmt.Sprintf("even: %s, odd: %s", evenResult, oddResult)
}

func main() {
	fmt.Println(ConcatenateStrings("-", "hello", "world", "how", "are", "you"))
}
