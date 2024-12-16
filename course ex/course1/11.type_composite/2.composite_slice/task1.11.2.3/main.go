package main

import "fmt"

func bitwiseXOR(n, res int) int {
	return n ^ res
}

func findSingleNumber(numbers []int) int {
	n := len(numbers)
	result := 0
	for i := 0; i < n; i++ {
		result = bitwiseXOR(result, numbers[i])
	}
	return result
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}

	singleNumber := findSingleNumber(numbers)
	fmt.Println(singleNumber)
}
