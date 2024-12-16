package main

import (
	"fmt"
)

func main() {
	xs := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(sum(xs))

	fmt.Println(average(xs))

	ys := [8]float64{1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5}

	fmt.Println(averageFloat(ys))

	fmt.Println(reverse(xs))
}

func sum(xs [8]int) int {
	result := 0

	for i := 0; i < len(xs); i++ {
		result += xs[i]
	}

	return result
}

func average(xs [8]int) float64 {
	result := 0.0
	for i := 0; i < len(xs); i++ {
		result += float64(xs[i])
	}
	return result / 8
}

func averageFloat(xs [8]float64) float64 {
	result := 0.0
	for i := 0; i < len(xs); i++ {
		result += xs[i]
	}
	return result / 8
}

func reverse(xs [8]int) [8]int {
	for i := 0; i < len(xs)/2; i++ {
		xs[i], xs[len(xs)-1-i] = xs[len(xs)-1-i], xs[i]
	}
	return xs
}
