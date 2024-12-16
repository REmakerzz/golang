package main

import "fmt"

type User struct {
	ID        int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}

func (u *User) TableName() string {
	return "users"
}

type Tabler interface {
	TableName() string
}

type SQLiteGenerator struct{}

type SQLGenerator interface {
	CreateTableSQL(table Tabler) string
	CreateInsertSQL(model Tabler) string
}

func (s *SQLiteGenerator) CreateTableSQL(table Tabler) string {
	panic("implement me")
}

func (s *SQLiteGenerator) CreateInsertSQL(model Tabler) string {
	panic("implement me")
}

type GoFakeitGenerator struct{}

type FakeDataGenerator interface {
	GenerateFakeUser() *User
}

func (g *GoFakeitGenerator) GenerateFakeUser() *User {
	panic("implement me")
}

func main() {
	fmt.Println("Program is running.")
}
