package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit"
)

type Person struct {
	Name string
	Age  int
}

func getUsers() []Person {
	var users []Person
	for i := 0; i < 10; i++ {
		user := Person{
			gofakeit.Name(), gofakeit.Number(18, 60),
		}
		users = append(users, user)
	}
	return users
}

func preparePrint(users []Person) string {
	var result string

	for _, user := range users {
		result += fmt.Sprintf("Имя: %s, Возраст: %d\n", user.Name, user.Age)
	}
	return result
}

func main() {
	users := getUsers()
	result := preparePrint(users)
	fmt.Println(result)
}
