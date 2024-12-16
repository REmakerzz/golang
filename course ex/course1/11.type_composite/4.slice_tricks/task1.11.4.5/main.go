package main

import (
	"fmt"
)

func main() {
	slice1 := []int{}
	slice2 := FilterDividers(slice1, 2)
	fmt.Println(slice2)
}

func FilterDividers(xs []int, divider int) []int {
	if divider == 0 {
		return xs
	}

	result := make([]int, 0, 0)
	for i := range xs {
		if xs[i]%divider == 0 {
			result = append(result, xs[i])
		}
	}
	return result
}
