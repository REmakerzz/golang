package main

import "fmt"

func concatStrings(xs ...string) string {
	var result string

	for _, word := range xs {
		result += word
	}
	return result
}

func main() {
	fmt.Println(concatStrings("Hello", " ", "word"))
}
