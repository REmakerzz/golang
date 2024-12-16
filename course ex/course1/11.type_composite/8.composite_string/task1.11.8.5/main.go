package main

import "fmt"

func ReverseString(str string) string {
	runes := []rune(str)

	lenght := len(runes)

	reversed := make([]rune, lenght)

	for i := 0; i < lenght; i++ {
		reversed[i] = runes[lenght-1-i]
	}

	return string(reversed)
}

func main() {
	fmt.Println(ReverseString("Hello, world!"))
	fmt.Println(ReverseString(""))
	fmt.Println(ReverseString("A"))
	fmt.Println(ReverseString("      "))
}
