package main

import (
	"fmt"
	"strings"
)

func main() {
	var num int = 3
	x := []int{3, 3, 4, 1, 2, 3, 5, 3, 2}
	y := 2
	s := "Was it a car or a cat I saw"
	name := "Ivan"
	fmt.Println(Factorial(&num))
	fmt.Println(isPalindrome(&s))
	fmt.Println(CountOccurrences(&x, &y))
	fmt.Println(ReverseString(&name))
}
//Factorial just factorial
func Factorial(n *int) int {
	if *n <= 1 {
		return 1
	}
	result := 1
	for i := 1; i <= *n; i++ {
		result *= i
	}
	return result
}

func isPalindrome(str *string) bool {
	s := strings.ToLower(strings.ReplaceAll(*str, " ", ""))
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
// CountOccurrences посчитать количество вхождения числа в срезе
func CountOccurrences(numbers *[]int, target *int) int {
	flag := 0
	for _, value := range *numbers {
		if value == *target {
			flag++
		}
	}
	return flag
}
// ReverseString just reverse string
func ReverseString(str *string) string {
	runes := []rune(*str)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}
