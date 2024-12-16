package main

import "fmt"

func main() {
	slice1 := []int{}
	slice2 := Cut(slice1, 0, 0)
	fmt.Println(slice2)
}

func Cut(xs []int, start, end int) []int {
	if start < 0 {
		return []int{}
	}

	if end < 0 {
		return []int{}
	}
	if start >= len(xs) {
		return []int{}
	}
	if end >= len(xs) {
		return []int{}
	}
	return xs[start : end+1]
}
