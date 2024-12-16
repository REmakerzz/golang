package main

import (
	"fmt"
	"strings"
)

func UserInfo(name string, age int, cities ...string) string {
	result := strings.Join(cities, ", ")
	return fmt.Sprintf("Имя: %s, возраст: %d, города: %s", name, age, result)
}

func main() {
	fmt.Println(UserInfo("Ivan", 25, "Moscow", "Питер", "Каза    нь"))
}
