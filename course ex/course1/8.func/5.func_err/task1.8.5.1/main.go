package main

import (
	"fmt"
	"strconv"
)

func CalculatePercentageChange(initialValue, finalValue string) (float64, error) {
	initialValueNum, err := strconv.ParseFloat(initialValue, 64)
	if err != nil {
		return 0, err
	}
	finalValueNum, err := strconv.ParseFloat(finalValue, 64)
	if err != nil {
		return 0, err
	}
	if initialValueNum == 0 {
		return 0, nil
	}
	return ((finalValueNum - initialValueNum) / initialValueNum) * 100, nil
}

func main() {
	a, err := CalculatePercentageChange("Hello", "World")
	fmt.Println(a, err)
}
