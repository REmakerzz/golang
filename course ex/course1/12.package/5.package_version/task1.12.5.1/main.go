package main

import (
	"fmt"

	"github.com/REmakerzz/mymath"
)

func Abs(x float64) float64 {
	return mymath.Abs(x)
}

func Max(x, y float64) float64 {
	return mymath.Max(x, y)
}

func Sqrt(x float64) float64 {
	return mymath.Sqrt(x)
}

func Yn(n int, x float64) float64 {
	return mymath.Yn(n, x)
}

func main() {
	fmt.Println(mymath.Max(2.0, 4.0))

}
