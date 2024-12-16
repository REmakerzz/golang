package main

import (
	"fmt"
)

func main() {
	num := []int{3, 5, 1, 8, 2}
	s := []string{"Hello", "world"}
	fmt.Println(Add(5, 10))
	fmt.Println(Max(num))
	fmt.Println(IsPrime(4))
	fmt.Println(ConcatenateStrings(s))
}
// Add sum
func Add(a, b int) *int {
	var res int
	var ptr *int = &res
	res = a + b
	return ptr
}
// Max max in slice
func Max(numbers []int) *int {
	max := numbers[0]
	var ptr *int
	for _, value := range numbers {
		if value > max {
			max = value
		}
	}
	ptr = &max
	return ptr
}
// IsPrime just isprime
func IsPrime(number int) *bool {
	var ptr *bool
	var result bool = true
	ptr = &result
	if number <= 1 {
		result = false
		ptr = &result
		return ptr
	}
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			result = false
			ptr = &result
			return ptr
		}
	}
	return ptr
}
// ConcatenateStrings just concatenate
func ConcatenateStrings(strs []string) *string {
	var ptr *string
	var result string

	for i := 0; i < len(strs); i++ {
		result += strs[i]
	}
	ptr = &result
	return ptr
}
