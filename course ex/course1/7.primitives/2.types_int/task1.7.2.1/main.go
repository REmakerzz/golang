package main

import "fmt"

func calculate(a int, b int) (int, int, int, int, int) {
	sum := a + b
	difference := a - b
	product := a * b
	quotient := a / b
	remainder := a % b
	return sum, difference, product, quotient, remainder
}


func main() {
    var sum int
    var difference int
    var product int
    var quotient int
    var remainder int
    //calculate(10, 3)
    sum, difference, product, quotient, remainder = calculate(10, 3)
    fmt.Printf("a = 10 b = 3 sum = %d difference = %d product = %d quotient = %d remainder = %d", sum, difference, product, quotient, remainder)
}