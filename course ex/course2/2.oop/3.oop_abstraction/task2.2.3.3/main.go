package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"reflect"
	"strings"
)

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

type SQLGenerator interface {
	CreateTableSQL(table Tabler) string
	CreateInsertSQL(model Tabler) string
}

type SQLiteGenerator struct{}

func (s *SQLiteGenerator) CreateTableSQL(table Tabler) string {
	var fields []string

	tType := reflect.TypeOf(table).Elem()
	for i := 0; i < tType.NumField(); i++ {
		field := tType.Field(i)
		dbField := field.Tag.Get("db_field")
		dbType := field.Tag.Get("db_type")
		if dbField == "" || dbType == "" {
			continue
		}
		fields = append(fields, fmt.Sprintf("%s %s", dbField, dbType))
	}
	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", table.TableName(), strings.Join(fields, ","))
}

func (s *SQLiteGenerator) CreateInsertSQL(model Tabler) string {
	tValue := reflect.ValueOf(model).Elem()
	tType := tValue.Type()

	var colums []string
	var values []string

	for i := 0; i < tType.NumField(); i++ {
		field := tType.Field(i)
		dbField := field.Tag.Get("db_field")
		if dbField == "" {
			continue
		}

		colums = append(colums, dbField)
		fieldValue := tValue.Field(i)
		if fieldValue.Kind() == reflect.String {
			values = append(values, fmt.Sprintf("'%s'", fieldValue.String()))
		} else {
			values = append(values, fmt.Sprintf("%v", fieldValue.Interface()))
		}
	}

	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);",
		model.TableName(),
		strings.Join(colums, ", "),
		strings.Join(values, ", "),
	)
}

type FakeDataGenerator interface {
	GenerateFakeUser() User
}

type GoFakeitGenerator struct{}

func (g *GoFakeitGenerator) GenerateFakeUser() User {
	return User{
		ID:        gofakeit.Number(1000, 9999),
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
	}
}

func GenerateUserInserts(count int) []string {
	var queries []string
	sqlGenerator := &SQLiteGenerator{}
	fakeDataGenerator := &GoFakeitGenerator{}

	for i := 0; i < count; i++ {
		user := fakeDataGenerator.GenerateFakeUser()
		query := sqlGenerator.CreateInsertSQL(&user)
		queries = append(queries, query)
	}

	return queries
}

func main() {
	gofakeit.Seed(0)

	sqlGenerator := &SQLiteGenerator{}
	fakeDataGenerator := &GoFakeitGenerator{}

	user := User{}
	sql := sqlGenerator.CreateTableSQL(&user)
	fmt.Println(sql)

	for i := 0; i < 34; i++ {
		fakeUser := fakeDataGenerator.GenerateFakeUser()
		query := sqlGenerator.CreateInsertSQL(&fakeUser)
		fmt.Println(query)
	}

	queries := GenerateUserInserts(34)
	for _, query := range queries {
		fmt.Println(query)
	}
}
