package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatalf("Ошибка при открытии базы данных: %v", err)
	}
	defer db.Close()

	err = CreateUserTable()
	if err != nil {
		log.Fatalf("Ошибка при создании таблицы: %v", err)
	}

	newUser := User{Name: "Arthur", Age: 30}
	err = InsertUser(newUser)
	if err != nil {
		log.Fatalf("Ошибка при добавлении пользователя: %v", err)
	}

	user, err := SelectUser(1)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("Выбранный пользователь: %+v\n", user)
	}

	user.Name = "Arthur Smith"
	user.Age = 31
	err = UpdateUser(user)
	if err != nil {
		log.Fatalf("Ошибка при обновлении пользователя: %v", err)
	} else {
		fmt.Printf("Обновленный пользователь: %v\n", user)
	}

	err = DeleteUser(1)
	if err != nil {
		log.Fatalf("Ошибка при удалении пользователя: %v", err)
	} else {
		fmt.Println("Удалил")
	}
}

func CreateUserTable() error {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatalf("Ошибка при открытии базы данных в функции CreateUserTable: %v", err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			age INTEGER NOT NULL)`)
	return err
}

func InsertUser(user User) error {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatalf("Ошибка при открытии базы данных в функции InsertUser: %v", err)
	}
	defer db.Close()

	_, err = db.Exec(`INSERT INTO users (name, age) VALUES (?, ?)`, user.Name, user.Age)
	return err
}

func SelectUser(id int) (User, error) {
	var user User

	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatalf("Ошибка при открытии базы данных в функции SelectUser: %v", err)
	}
	defer db.Close()

	row := db.QueryRow(`SELECT id, name, age FROM users WHERE id = ?`, id)
	err = row.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, fmt.Errorf("пользователь с ID %d не найден", id)
		}
		return User{}, err
	}
	return user, nil
}

func UpdateUser(user User) error {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatalf("Ошибка при открытии базы данных в функции UpdateUser: %v", err)
	}
	defer db.Close()

	_, err = db.Exec(`UPDATE users SET name = ?, age = ? WHERE id = ?`, user.Name, user.Age, user.ID)
	return err
}

func DeleteUser(id int) error {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatalf("Ошибка при открытии базы данных в фукнции DeleteUser: %v", err)
	}
	defer db.Close()

	_, err = db.Exec(`DELETE FROM users WHERE id = ?`, id)
	return err
}
