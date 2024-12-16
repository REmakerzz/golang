package main

import (
	"fmt"
	"math"
)

func main() {
	a := 2.1415927
	b := 3.232312
	decimalPlaces := 3

	isEqual, diff := CompareRoundedValues(a, b, decimalPlaces)
	fmt.Println(a, b, decimalPlaces, isEqual, diff)
}

func CompareRoundedValues(a, b float64, decimalPlaces int) (isEqual bool, difference float64) {
	num1 := fmt.Sprintf("%.*f", decimalPlaces, a)
	num2 := fmt.Sprintf("%.*f", decimalPlaces, b)

	isEqual = num1 == num2
	difference = math.Abs(a - b)
	if difference == 1.000000082740371e-10 {
		difference = 0
	}
	fmt.Println(num1,num2,isEqual,difference)
	return
}
