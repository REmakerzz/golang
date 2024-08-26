package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Словарь для конвертации римских чисел в арабские
var romanToArabic = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

// Массив для конвертации арабских чисел в римские
var arabicToRoman = []string{
	"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
}

// Функция для конвертации римского числа в арабское
func romanToInt(roman string) (int, error) {
	if val, ok := romanToArabic[roman]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("invalid Roman numeral")
}

// Функция для конвертации арабского числа в римское
func intToRoman(num int) (string, error) {
	if num > 0 && num < len(arabicToRoman) {
		return arabicToRoman[num], nil
	}
	return "", fmt.Errorf("result out of Roman numeral range")
}

// Функция для выполнения арифметической операции
func calculate(a int, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operation")
	}
}

// Функция проверки диапазона чисел
func validateInputRange(num int) {
	if num < 1 || num > 10 {
		panic("input number out of range (must be between 1 and 10)")
	}
}

// Главная функция
func main() {
	fmt.Println("Введите выражение:")
	var input string
	fmt.Scanln(&input)

	// Определение оператора и разделение строки на операнды
	var operator string
	if strings.Contains(input, "+") {
		operator = "+"
	} else if strings.Contains(input, "-") {
		operator = "-"
	} else if strings.Contains(input, "*") {
		operator = "*"
	} else if strings.Contains(input, "/") {
		operator = "/"
	} else {
		panic("invalid operator")
	}

	parts := strings.Split(input, operator)
	if len(parts) != 2 {
		panic("invalid input format")
	}

	left := strings.TrimSpace(parts[0])
	right := strings.TrimSpace(parts[1])

	// Обработка римских чисел
	isRoman := false
	a, err1 := strconv.Atoi(left)
	b, err2 := strconv.Atoi(right)

	if err1 != nil || err2 != nil {
		isRoman = true
		a, err1 = romanToInt(left)
		b, err2 = romanToInt(right)
		if err1 != nil || err2 != nil {
			panic("invalid Roman numerals")
		}
	} else {
		// Проверка диапазона для арабских чисел
		validateInputRange(a)
		validateInputRange(b)
	}

	// Выполнение арифметической операции
	result := calculate(a, b, operator)

	// Вывод результата
	if isRoman {
		if result < 1 {
			panic("Roman numeral result less than I")
		}
		romanResult, _ := intToRoman(result)
		fmt.Println("Результат:", romanResult)
	} else {
		fmt.Println("Результат:", result)
	}
}
