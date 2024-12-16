package main

import (
	"fmt"
)

func CheckDiscount(price, discount float64) (float64, error) {
	if discount > 50 {
		return 0, fmt.Errorf("Скидка не может превышать 50%%")
	}
	discountedPrice := price * (1 - discount/100)
	return discountedPrice, nil
}

func main() {
	a, err := CheckDiscount(25333, 25)
	fmt.Println(a, err)
}
