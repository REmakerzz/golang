package main

import (
	"fmt"
	"math"
)

func main() {
	resultSin := Sin(45)
	resultCos := Cos(45)
	fmt.Println(resultSin, resultCos)
}

func Sin(x float64) float64 {
	return math.Sin(x)
}

func Cos(x float64) float64 {
	return math.Cos(x)
}
