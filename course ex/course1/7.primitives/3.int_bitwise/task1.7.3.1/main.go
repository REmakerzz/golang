package main

import "fmt"

func main() {
	fmt.Println("a & b =", bitwiseAnd(43, 2))
	fmt.Println("a | b =", bitwiseOr(43, 2))
	fmt.Println("a ^ b =", bitwiseXor(43, 2))
	fmt.Println("a << b =", bitwiseLeftShift(43, 2))
	fmt.Println("a >> b =", bitwiseRightShift(43, 2))
}

func bitwiseAnd(a, b int) int {
	return a & b
}

func bitwiseOr(a, b int) int {
	return a | b
}

func bitwiseXor(a, b int) int {
	return a ^ b
}

func bitwiseLeftShift(a, b int) int {
	return a << b
}

func bitwiseRightShift(a, b int) int {
	return a >> b
}
