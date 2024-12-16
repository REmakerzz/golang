package main

import "fmt"

func main() {

	slice1 := []int{1, 2, 3}
	slice2 := InsertToStart(slice1, 4, 5, 6)
	fmt.Println(slice2)
}

func InsertToStart(xs []int, x ...int) []int {
	if len(x) <= 0 {
		return xs
	}

	result := make([]int, 0, len(x))
	for i := range x {
		result = append(result, x[i])
	}
	result = append(result, xs...)
	return result
}
