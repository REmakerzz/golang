package main

import "fmt"

func UserInfo(name, city, phone string, age, weight int) string {
	return fmt.Sprintf("Имя: %s, Город: %s, Телефон: %s, Возраст: %d, Вес: %d", name, city, phone, age, weight)
}

func main() {
	fmt.Println(UserInfo("Ivan", "Moscow", "+7(999)999 99 99", 22, 75))
}
