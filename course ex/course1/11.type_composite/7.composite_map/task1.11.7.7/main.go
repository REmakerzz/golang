package main

import "fmt"

type User struct {
	Nickname string
	Age      int
	Email    string
}

func getUniqueUsers(users []User) []User {
	uniqueMap := make(map[string]struct{})

	uniqueUsers := []User{}

	for _, user := range users {
		if _, exists := uniqueMap[user.Nickname]; !exists {
			uniqueMap[user.Nickname] = struct{}{}
			uniqueUsers = append(uniqueUsers, user)
		}
	}

	return uniqueUsers
}

func main() {
	users := []User{
		{Nickname: "Alice", Age: 30, Email: "alice@example.com"},
		{Nickname: "Bob", Age: 25, Email: "bob@example.com"},
	}

	uniqueUsers := getUniqueUsers(users)
	fmt.Println(uniqueUsers)
}
