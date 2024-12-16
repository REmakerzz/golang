package main

import "fmt"

func appendInt(xs []int, x ...int) []int {
	return append(xs, x...)
}

func main() {
	slice1 := []int{1, 2, 3}

	slice2 := appendInt(slice1, 4, 5, 6)
	fmt.Println(slice1, slice2)
}
