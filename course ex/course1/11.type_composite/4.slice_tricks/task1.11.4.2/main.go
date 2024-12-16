package main

import "fmt"

func main() {
	xs := []int{1, 2, 3, 4, 5}
	result := InsertAfterIDX(xs, 2, 6, 7, 8)
	fmt.Println(result)
}

func InsertAfterIDX(xs []int, idx int, x ...int) []int {
	if idx < 0 || idx >= len(xs) {
		return xs
	}
	n := len(xs) + len(x)

	result := make([]int, 0, n)

	result = append(result, xs[:idx+1]...)
	result = append(result, x...)
	result = append(result, xs[idx+1:]...)
	return result
}
