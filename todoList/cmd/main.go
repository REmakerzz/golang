package main

import (
	"log"
	"net/http"
	"todoList/internal/database"
	"todoList/internal/handlers"
)

func main() {
	// Подключаемся к базе данных
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Создаём маршруты
	http.HandleFunc("/tasks", handlers.TasksHandler(db))
	http.HandleFunc("/task/", handlers.TaskByIDHandler(db))

	// Запускаем сервер
	log.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
