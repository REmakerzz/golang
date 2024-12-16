package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID       int
	Username string
	Email    string
}

func CreateUserTable() error {

	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE
		)`
	_, err = db.Exec(createTableSQL)
	return err
}

func InsertUser(user User) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query, args, err := PrepareQuery("insert", "users", user)
	if err != nil {
		return err
	}

	_, err = db.Exec(query, args...)
	return err
}

func SelectUser(userID int) (User, error) {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	query, args, err := PrepareQuery("select", "users", User{ID: userID})
	if err != nil {
		return User{}, err
	}

	row := db.QueryRow(query, args...)
	var user User
	err = row.Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func UpdateUser(user User) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query, args, err := PrepareQuery("update", "users", user)
	if err != nil {
		return err
	}

	_, err = db.Exec(query, args...)
	return err
}

func DeleteUser(userID int) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query, args, err := PrepareQuery("delete", "users", User{ID: userID})
	if err != nil {
		return err
	}

	_, err = db.Exec(query, args...)

	return err
}

func PrepareQuery(operation string, table string, user User) (string, []interface{}, error) {
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question)

	switch operation {
	case "insert":
		query, args, err := psql.Insert(table).Columns("username", "email").Values(user.Username, user.Email).ToSql()
		return query, args, err

	case "select":
		query, args, err := psql.Select("id", "username", "email").From(table).Where("id = ?", user.ID).ToSql()
		return query, args, err

	case "update":
		query, args, err := psql.Update(table).Set("username", user.Username).Set("email", user.Email).Where("id = ?", user.ID).ToSql()
		return query, args, err

	case "delete":
		query, args, err := psql.Delete(table).Where("id = ?", user.ID).ToSql()
		return query, args, err

	default:
		return "", nil, fmt.Errorf("unsupported operation: %s", operation)
	}
}

func main() {
	err := CreateUserTable()
	if err != nil {
		log.Fatalf("Ошибка при создании таблицы: %v", err)
	}

	user := User{Username: "Arthur_Smith", Email: "arthur@example.com"}
	err = InsertUser(user)
	if err != nil {
		log.Fatalf("Ошибка при вставке пользователя: %v", err)
	} else {
		fmt.Println("Пользователь успешно добавлен.")
	}

	selectedUser, err := SelectUser(1)
	if err != nil {
		log.Fatalf("Ошибка при выборке пользователя: %v", err)
	} else {
		fmt.Printf("Выбранный Пользователь: %+v\n", selectedUser)
	}

	selectedUser.Email = "arthur_smith@example.com"
	err = UpdateUser(selectedUser)
	if err != nil {
		log.Fatalf("Ошибка при обновлении пользователя: %v", err)
	} else {
		fmt.Println("Информация о пользователе обновлена.")
	}

	err = DeleteUser(1)
	if err != nil {
		log.Fatalf("Ошибка при удалении пользователя: %v", err)
	} else {
		fmt.Println("Пользователь удален.")
	}
}
