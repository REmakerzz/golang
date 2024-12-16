package main

import "fmt"

func CalculatePercentageChange(initialValue, finalValue float64) float64 {
	if initialValue == 0 {
		return 0
	}
	return ((finalValue - initialValue) / initialValue) * 100
}

func main() {
	var a float64
	a = CalculatePercentageChange(10, 50)
	fmt.Println(a)
}
