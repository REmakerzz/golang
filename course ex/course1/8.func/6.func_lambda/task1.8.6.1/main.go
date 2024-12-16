package main

import "fmt"

func Sum(a ...int) int {
	var sum int = 0
	for _, value := range a {
		sum += value
	}
	return sum
}

func Mul(a ...int) int {
	var result int = 1
	for _, num := range a {
		result *= num
	}
	return result
}

func Sub(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	result := a[0]
	for _, num := range a[1:] {
		result -= num
	}
	return result
}

func MathOperate(op func(a ...int) int, a ...int) int {
	return op(a...)
}

func main() {
	fmt.Println(MathOperate(Sum, 1, 1, 3))
	fmt.Println(MathOperate(Mul, 1, 7, 3))
	fmt.Println(MathOperate(Sub, 13, 2, 3))
}
