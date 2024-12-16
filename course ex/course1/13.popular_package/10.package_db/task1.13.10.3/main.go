package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	UserID int    `json:"user_id"`
}

func CreateUserTable() error {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER NOT NULL)`)

	if err != nil {

		return err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS comments (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	text TEXT NOT NULL,
    	user_id INTEGER NOT NULL,
    	FOREIGN KEY(user_id) REFERENCES users(id))`)
	if err != nil {
		return err
	}

	fmt.Println("Tables created successfully")

	return nil
}

func InsertUser(user User) error {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		return err
	}

	defer db.Close()

	query, args, err := prepareQuery("insert", "users", user).(squirrel.InsertBuilder).ToSql()
	if err != nil {
		return err
	}
	result, err := db.Exec(query, args...)
	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	for _, comment := range user.Comments {
		comment.UserID = int(userID)
		err := InsertComment(db, comment)
		if err != nil {
			return err
		}
	}
	return nil
}

func InsertComment(db *sql.DB, comment Comment) error {
	query, args, err := prepareQuery("insert", "comments", comment).(squirrel.InsertBuilder).ToSql()
	if err != nil {
		return err
	}
	_, err = db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func SelectUser(userID int) (User, error) {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	var user User
	query, args, err := prepareQuery("select", "users", User{ID: userID}).(squirrel.SelectBuilder).ToSql()
	if err != nil {
		return User{}, err
	}
	row := db.QueryRow(query, args...)
	err = row.Scan(&user.ID, &user.Name, &user.Age)
	if err == sql.ErrNoRows {
		return User{}, nil
	} else if err != nil {
		return User{}, err
	}

	user.Comments, err = SelectCommentsByUserID(user.ID)
	return user, nil
}

func SelectCommentsByUserID(userID int) ([]Comment, error) {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var comments []Comment
	query, args, err := prepareQuery("select", "comments", Comment{UserID: userID}).(squirrel.SelectBuilder).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.ID, &comment.Text, &comment.UserID); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func UpdateUser(user User) error {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query, args, err := prepareQuery("update", "users", user).(squirrel.UpdateBuilder).ToSql()
	if err != nil {
		return err
	}
	_, err = db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(userID int) error {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query, args, err := prepareQuery("delete", "users", User{ID: userID}).(squirrel.DeleteBuilder).ToSql()
	if err != nil {
		return err
	}
	_, err = db.Exec(query, args...)
	if err != nil {
		return err
	}

	tx, _ := db.Begin()
	defer tx.Rollback()

	if _, err := tx.Exec(query, args...); err != nil {
		return err
	}
	if _, err := tx.Exec(`DELETE FROM comments WHERE user_id = ?`, userID); err != nil {
		return err
	}
	return tx.Commit()
}

func prepareQuery(operation string, table string, user interface{}) interface{} {
	switch operation {
	case "insert":
		switch t := user.(type) {
		case User:
			return squirrel.Insert(table).Columns("name", "age").Values(t.Name, t.Age)
		case Comment:
			return squirrel.Insert("comments").Columns("text", "user_id").Values(t.Text, t.UserID)
		}
	case "select":
		switch t := user.(type) {
		case User:
			return squirrel.Select("*").From(table).Where(squirrel.Eq{"id": t.ID})
		case Comment:
			return squirrel.Select("*").From("comments").Where(squirrel.Eq{"user_id": t.UserID})
		}
	case "update":
		return squirrel.Update(table).Set("name", user.(User).Name).Set("age", user.(User).Age).Where(squirrel.Eq{"id": user.(User).ID})
	case "delete":
		return squirrel.Delete(table).Where(squirrel.Eq{"id": user.(User).ID})
	default:
		return nil
	}
	return nil
}

func main() {
	err := CreateUserTable()
	if err != nil {
		log.Fatalf("Ошибка при создании таблицы: %v", err)
	}

	comments := []Comment{{Text: "Comment_1", UserID: 0}, {Text: "Comment_2", UserID: 0}}
	newUser := User{Name: "Arthur", Age: 30, Comments: comments}
	err = InsertUser(newUser)
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

	selectedUser.Name = "Arthur Smith"
	err = UpdateUser(selectedUser)
	if err != nil {
		log.Fatalf("Ошибка при обновлении пользователя: %v", err)
	} else {
		fmt.Printf("Информация о пользователе обновлена: %+v\n", selectedUser)
	}

	err = DeleteUser(1)
	if err != nil {
		log.Fatalf("Ошибка при удалении пользователя: %v", err)
	} else {
		fmt.Println("Пользователь удален.")
	}
}
