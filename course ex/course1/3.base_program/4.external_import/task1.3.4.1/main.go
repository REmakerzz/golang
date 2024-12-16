package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	fmt.Println(DecimalSum("4.5", "2.4"))
	fmt.Println(DecimalSubtract("6.4", "3.2"))
	fmt.Println(DecimalMultiply("2.3", "7.4"))
	fmt.Println(DecimalDivide("4.2", "2.2"))
	fmt.Println(DecimalRound("4.423", 1))
	fmt.Println(DecimalGreaterThan("4.2", "2.3"))
	fmt.Println(DecimalLessThan("7.7", "6.4"))
	fmt.Println(DecimalEqual("4.4", "4.4"))
}
// DecimalSum calculates the sum of two decimal numbers and returns the result.
func DecimalSum(a, b string) (string, error) {
	x, err := decimal.NewFromString(a)
	if err != nil {
		panic(err)
	}
	y, err := decimal.NewFromString(b)
	if err != nil {
		panic(err)
	}
	sum := decimal.Sum(x, y)
	return sum.String(), err
}
// DecimalSubtract calculates the sub of two decimal numbers and returns the result.
func DecimalSubtract(a, b string) (string, error) {
	x, err := decimal.NewFromString(a)
	if err != nil {
		panic(err)
	}
	y, err := decimal.NewFromString(b)
	if err != nil {
		panic(err)
	}
	MaxNum := decimal.Max(x, y)
	MinNum := decimal.Min(x, y)
	result := MaxNum.Sub(MinNum)
	return result.String(), err
}
// DecimalMultiply которая принимает две строки, содержащие числа с плавающей точкой,
//перемножает их и возвращает результат в виде строки. Если входные данные некорректны, функция должна возвращать ошибку.
func DecimalMultiply(a, b string) (string, error) {
	x, err := decimal.NewFromString(a)
	if err != nil {
		panic(err)
	}
	y, err := decimal.NewFromString(b)
	if err != nil {
		panic(err)
	}
	result := x.Mul(y)
	return result.String(), err
}
// DecimalDivide функция, которая принимает две строки, содержащие числа с плавающей точкой,
// делит первое число на второе и возвращает результат в виде строки. 
// Если входные данные некорректны или происходит на ноль, функция должна возвращать ошибку.
func DecimalDivide(a, b string) (string, error) {
	x, err := decimal.NewFromString(a)
	if err != nil {
		panic(err)
	}
	y, err := decimal.NewFromString(b)
	if err != nil {
		panic(err)
	}
	if y.IsZero() {
		err := "Деление на ноль"
		panic(err)
	}
	result := x.Div(y)
	return result.String(), err
}
// DecimalRound функция, которая принимает строку, содержащую число с плавающей точкой,
// и точность округления в виде int32. Функция должна округлить число до указанной точности
// и вернуть результат в виде строки. Если входные данные некорректны, функция должна возвращать ошибку.
func DecimalRound(a string, precision int32) (string, error) {
	x, err := decimal.NewFromString(a)
	if err != nil {
		panic(err)
	}
	result := x.Round(precision)
	return result.String(), err
}
// DecimalGreaterThan функция, которая принимает две строки, содержащие числа с плавающей точкой,
// и возвращает true, если первое число больше второго, и false в противном случае.
// Если входные данные некорректны, функция должна возвращать ошибку.
func DecimalGreaterThan(a, b string) (bool, error) {
	x, err := decimal.NewFromString(a)
	if err != nil {
		panic(err)
	}
	y, err := decimal.NewFromString(b)
	if err != nil {
		panic(err)
	}
	result := x.GreaterThan(y)
	return result, err
}
// DecimalLessThan функция, которая принимает две строки, содержащие числа с плавающей точкой,
// и возвращает true, если первое число меньше второго, и false в противном случае.
// Если входные данные некорректны, функция должна возвращать ошибку.
func DecimalLessThan(a, b string) (bool, error) {
	x, err := decimal.NewFromString(a)
	if err != nil {
		panic(err)
	}
	y, err := decimal.NewFromString(b)
	if err != nil {
		panic(err)
	}
	result := x.LessThan(y)
	return result, err
}
// DecimalEqual функция, которая принимает две строки, содержащие числа с плавающей точкой,
// и возвращает true, если числа равны, и false в противном случае.
// Если входные данные некорректны, функция должна возвращать ошибку.
func DecimalEqual(a, b string) (bool, error) {
	x, err := decimal.NewFromString(a)
	if err != nil {
		panic(err)
	}
	y, err := decimal.NewFromString(b)
	if err != nil {
		panic(err)
	}
	result := x.Equal(y)
	return result, err
}
