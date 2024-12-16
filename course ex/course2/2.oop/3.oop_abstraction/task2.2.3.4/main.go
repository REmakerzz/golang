package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
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

func (s *SQLiteGenerator) CreateTableSQL(t Tabler) string {
	tType := reflect.TypeOf(t).Elem()
	var columns []string

	for i := 0; i < tType.NumField(); i++ {
		field := tType.Field(i)
		dbField := field.Tag.Get("db_field")
		dbType := field.Tag.Get("db_type")
		if dbField != "" && dbType != "" {
			columns = append(columns, fmt.Sprintf("%s %s", dbField, dbType))
		}
	}

	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", t.TableName(), strings.Join(columns, ","))
}

func (s *SQLiteGenerator) CreateInsertSQL(t Tabler) string {
	tValue := reflect.ValueOf(t).Elem()
	tType := tValue.Type()

	var columns []string
	var values []string

	for i := 0; i < tValue.NumField(); i++ {
		field := tType.Field(i)
		dbField := field.Tag.Get("db_field")
		if dbField == "" {
			continue
		}
		columns = append(columns, dbField)

		fieldValue := tValue.Field(i)
		if fieldValue.Kind() == reflect.String {
			values = append(values, fmt.Sprintf("'%s'", fieldValue.String()))
		} else {
			values = append(values, fmt.Sprintf("%v", fieldValue.Interface()))
		}
	}

	return fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s);",
		t.TableName(),
		strings.Join(columns, ","),
		strings.Join(values, ","),
	)
}

type Migrator struct {
	db           *sql.DB
	sqlGenerator SQLGenerator
}

func NewMigrator(db *sql.DB, sqlGenerator SQLGenerator) *Migrator {
	return &Migrator{
		db:           db,
		sqlGenerator: sqlGenerator,
	}
}

func (m *Migrator) Migrate(models ...Tabler) error {
	for _, model := range models {
		createTableSQL := m.sqlGenerator.CreateTableSQL(model)
		_, err := m.db.Exec(createTableSQL)
		if err != nil {
			return fmt.Errorf("failed to create table for model %s: %w", model.TableName(), err)
		}
		log.Printf("Table %s created successfully", model.TableName())
	}
	return nil
}

func main() {
	db, err := sql.Open("sqlite3", "./my_database.db")
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer db.Close()

	sqlGenerator := &SQLiteGenerator{}

	migrator := NewMigrator(db, sqlGenerator)

	if err := migrator.Migrate(&User{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}
}
