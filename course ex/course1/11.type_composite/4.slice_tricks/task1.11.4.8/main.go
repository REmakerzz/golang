package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4}
	firstElement, shiftedSlice := Shift(slice1)
	fmt.Println(firstElement)
	fmt.Println(shiftedSlice)
}

func Shift(xs []int) (int, []int) {
	//Проверка на пустой слайс
	if len(xs) == 0 {
		return 0, []int{}
	}

	//создаю слайс результата фукнции с длинной оригинального слайса
	result := make([]int, len(xs))
	// первому элементу слайса результата фукнции присваиваю последний элемент оригинального слайса
	result[0] = xs[len(xs)-1]
	// копирую элементы оригинального слайса в слайс результата функции
	copy(result[1:], xs)
	return xs[0], result
}
