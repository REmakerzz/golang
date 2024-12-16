package main

import "fmt"

func DivideAndRemainder(a, b int) (int, int) {
	defer func() {
		if r := recover(); r != nil {
			if b == 0 {
				fmt.Println("Деление на ноль")
			}
			fmt.Println("check zero argument")
		}
	}()
	return a / b, a % b
}
func main() {
	var a, b int
	a, b = DivideAndRemainder(10, 2)
	fmt.Printf("Частное: %d, Остаток: %d", a, b)
}
