package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}

	appendInt(&slice1, 4, 5, 6)
	fmt.Println(slice1)
}

func appendInt(xs *[]int, x ...int) {
	*xs = append(*xs, x...)
}
