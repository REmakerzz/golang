package main

import (
	"testing"
)

func expectPanic(t *testing.T, f func(), expectedMessage string) {
	defer func() {
		if r := recover(); r != nil {
			if r != expectedMessage {
				t.Errorf("expected panic message %q, got %q", expectedMessage, r)
			}
		} else {
			t.Errorf("expected panic, but no panic occurred")
		}
	}()
	f()
}

func TestSQLiteGenerator_CreateTableSQL(t *testing.T) {
	sqlGenerator := &SQLiteGenerator{}
	expectPanic(t, func() {
		sqlGenerator.CreateTableSQL(nil)
	}, "implement me")
}

func TestSQLiteGenerator_CreateInsertSQL(t *testing.T) {
	sqlGenerator := &SQLiteGenerator{}
	expectPanic(t, func() {
		sqlGenerator.CreateInsertSQL(nil)
	}, "implement me")
}

func TestGoFakeitGenerator_GenerateFakeUser(t *testing.T) {
	fakeDataGenerator := &GoFakeitGenerator{}
	expectPanic(t, func() {
		fakeDataGenerator.GenerateFakeUser()
	}, "implement me")
}
