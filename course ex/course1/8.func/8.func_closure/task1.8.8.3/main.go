package main

import "fmt"

func adder(initial int) func(int) int {
	return func(y int) int {
		return y + initial
	}
}

func main() {
	addTwo := adder(2)
	result := addTwo(3)
	fmt.Println(result)
}
