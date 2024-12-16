package main

import (
	"fmt"
	"math"
)

var CalculateCircleArea = func(radius float64) float64 {
	result := math.Pi * math.Pow(radius, 2)
	return result
}

var CalculateRectangleArea = func(width, height float64) float64 {
	return width * height
}

var CalculateTriangleArea = func(base, height float64) float64 {
	return 0.5 * (base * height)
}

func main() {
	fmt.Println(CalculateCircleArea(2))
	fmt.Println(CalculateRectangleArea(10, 10))
	fmt.Println(CalculateTriangleArea(2, 10))
}
