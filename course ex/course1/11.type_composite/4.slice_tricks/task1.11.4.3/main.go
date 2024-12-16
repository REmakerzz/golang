package main

import "fmt"

func main() {
	slice1 := make([]int, 40, 100)
	slice1 = append(slice1, 1, 2, 3, 4, 5, 6)
	slice2 := RemoveExtraMemory(slice1)
	fmt.Println(slice1, slice2, len(slice1), len(slice2), cap(slice1), cap(slice2))
}

func RemoveExtraMemory(xs []int) []int {
	if len(xs) == cap(xs) {
		return xs
	}

	result := make([]int, len(xs), len(xs))
	copy(result, xs)

	return result
}
