package main

import "fmt"

func MaxDifference(numbers []int) int {
	if len(numbers) == 1 {
		return 0
	}
	if len(numbers) == 0 {
		return 0
	}
	n := len(numbers)
	maxValue, minValue := numbers[n-1], numbers[0]
	for i := 0; i < n; i++ {
		if numbers[i] < minValue {
			minValue = numbers[i]
		}
		if numbers[i] > maxValue {
			maxValue = numbers[i]
		}
	}
	return maxValue - minValue
}

func main() {
	xs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := MaxDifference(xs)
	fmt.Println(result)
}
