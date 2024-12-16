package main

import "fmt"

func main() {
	var x int = 22
	var s string = "fooboobar"
	mutate(&x)
	fmt.Println(x)
	ReverseString(&s)
	fmt.Println(s)
}

func mutate(a *int) {
	*a = 42
}
// ReverseString just reverse string
func ReverseString(str *string) {
	runes := []rune(*str)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	*str = string(runes)
}
