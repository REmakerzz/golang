package main

import (
	"fmt"
	"math"
)

func main() {
	result := Sqrt(4)
	fmt.Println(result)
}

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}
