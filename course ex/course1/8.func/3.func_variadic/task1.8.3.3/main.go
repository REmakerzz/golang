package main

import "fmt"

func PrintNumbers(numbers ...int) {
	for _, num := range numbers {
		fmt.Println(num)
	}
}

func main() {
	PrintNumbers(1, 2, 3, 4, 5)
}
