package main

import (
	"encoding/json"
	"fmt"
)

type Comment struct {
	Text string `json:"text"`
}

type User struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}

func getJSON(data []User) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("error Marshal json: %w", err)
	}
	return string(jsonData), nil
}

func main() {

	users := []User{
		{
			Name: "Alice",
			Age:  30,
			Comments: []Comment{
				{Text: "Great post!"},
				{Text: "Thanks for sharing."},
			},
		},
		{
			Name: "Bob",
			Age:  25,
			Comments: []Comment{
				{Text: "Very informative."},
			},
		},
	}

	jsonStr, err := getJSON(users)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("JSON string:", jsonStr)
}
