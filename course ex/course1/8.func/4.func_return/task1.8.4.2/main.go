package main

import "fmt"

func FindMaxAndMin(n ...int) (int, int) {
	if len(n) == 0 {
		return 0, 0
	}
	maxValue := n[0]
	minValue := n[0]
	for i := 0; i < len(n); i++ {
		if maxValue < n[i] {
			maxValue = n[i]
		}
		if minValue > n[i] {
			minValue = n[i]
		}
	}
	return maxValue, minValue
}

func main() {
	var a, b int
	a, b = FindMaxAndMin()
	fmt.Println(a, b)
}
