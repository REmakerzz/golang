package main

import (
	"fmt"
	"math"
)

func main() {
	result := Floor(3.6)
	fmt.Println(result)
}

func Floor(x float64) float64 {
	return math.Floor(x)
}
