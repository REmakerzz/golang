package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice2 := RemoveIDX(slice1, 2)
	fmt.Println(slice2)
}

func RemoveIDX(xs []int, idx int) []int {
	if idx > len(xs) {
		return xs
	}
	if idx < 0 {
		return xs
	}

	result := make([]int, 0, len(xs))
	result = append(result, xs[:idx]...)
	result = append(result, xs[idx+1:]...)
	return result
}
