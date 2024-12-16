package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	firstIDX, slice2 := Pop(slice1)
	fmt.Println(firstIDX, slice2)
}

func Pop(xs []int) (int, []int) {
	if len(xs) <= 0 {
		return 0, []int{}
	}

	return xs[0], xs[1:]
}
