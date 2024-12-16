package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(hypotenuse(3.4,1.2))
}

func hypotenuse(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
 }