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

type Tabler interface {
	TableName() string
}

func (u *User) TableName() string {
	return "users"
}

type SQLGenerator interface {
	CreateTableSQL(table Tabler) string
	CreateInsertSQL(model Tabler) string
}

type SQLiteGenerator struct{}

func (s *SQLiteGenerator) CreateTableSQL(table Tabler) string {
	t := reflect.TypeOf(table)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		panic("CreateTableSQL expects a struct type")
	}

	var fields []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		dbField := field.Tag.Get("db_field")
		dbType := field.Tag.Get("db_type")
		fields = append(fields, fmt.Sprintf("%s %s", dbField, dbType))
	}

	return fmt.Sprintf("CREATE TABLE %s (\n %s\n);", table.TableName(), strings.Join(fields, ",\n "))
}

func (s *SQLiteGenerator) CreateInsertSQL(model Tabler) string {
	v := reflect.ValueOf(model)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic("CreateInsertSQL expects a struct type")
	}

	typ := v.Type()

	var columns, values []string
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		dbField := field.Tag.Get("db_field")
		columns = append(columns, dbField)
		values = append(values, fmt.Sprintf("'%v'", v.Field(i).Interface()))
	}

	return fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s);",
		model.(Tabler).TableName(),
		strings.Join(columns, ", "),
		strings.Join(values, ", "),
	)
}

type FakeDataGenerator interface {
	GenerateFakeUser() User
}

type GoFakeitGenerator struct {
	currentID int
}

func (g *GoFakeitGenerator) GenerateFakeUser() *User {
	g.currentID++
	return &User{
		ID:        g.currentID,
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
	}
}

func main() {
	sqlGenerator := &SQLiteGenerator{}
	fakeDataGenerator := &GoFakeitGenerator{}

	user := &User{}
	sql := sqlGenerator.CreateTableSQL(user)
	fmt.Println(sql)

	for i := 0; i < 34; i++ {
		fakeUser := fakeDataGenerator.GenerateFakeUser()
		query := sqlGenerator.CreateInsertSQL(fakeUser)
		fmt.Println(query)
	}
}
