package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"todoList/internal/models"
)

// Получение всех задач
func TasksHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			rows, err := db.Query("SELECT id, name, done FROM tasks")
			if err != nil {
				http.Error(w, "Failed to query tasks", http.StatusInternalServerError)
				return
			}
			defer rows.Close()

			var tasks []models.Task
			for rows.Next() {
				var task models.Task
				if err := rows.Scan(&task.ID, &task.Name, &task.Done); err != nil {
					http.Error(w, "Failed to scan task", http.StatusInternalServerError)
					return
				}
				tasks = append(tasks, task)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(tasks)
		}
	}
}

// Работа с задачей по ID
func TaskByIDHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/task/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid task ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			var task models.Task
			err := db.QueryRow("SELECT id, name, done FROM tasks WHERE id = $1", id).Scan(&task.ID, &task.Name, &task.Done)
			if err == sql.ErrNoRows {
				http.Error(w, "Task not found", http.StatusNotFound)
				return
			} else if err != nil {
				http.Error(w, "Failed to query task", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)

		case http.MethodDelete:
			_, err := db.Exec("DELETE FROM tasks WHERE id = $1", id)
			if err != nil {
				http.Error(w, "Failed to delete task", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusNoContent)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
