package main

import "fmt"

func CalculateStockValue(price float64, quantity int) (float64, float64) {
	total := price * float64(quantity)
	return total, price
}

func main() {
	var a, b float64
	a, b = CalculateStockValue(12.2, 2)
	fmt.Println(a, b)
}
