package main

import "fmt"

func getBytes(s string) []byte {
	return []byte(s)
}

func getRunes(s string) []rune {
	return []rune(s)
}

func main() {
	fmt.Println(getBytes("Привет, мир!"))

	fmt.Println(getRunes("Привет, мир!"))
}
