package main

import (
	"fmt"

	"github.com/REmakerzz/mymath"
)

func Sqrt(x float64) float64 {
	return mymath.Sqrt(x)
}

func Abs(x float64) float64 {
	return mymath.Abs(x)
}

func Max(x float64, y float64) float64 {
	return mymath.Max(x, y)
}

func Yn(n int, x float64) float64 {
	return mymath.Yn(n, x)
}

func main() {
	fmt.Println(Sqrt(4))

	fmt.Println(Max(2.0, 4.0))
	fmt.Println(Yn(1, 2.0))
	fmt.Println(Abs(2.0))
}
