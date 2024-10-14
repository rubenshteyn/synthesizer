package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"synthesizer/pkg/models"
)

func GetTasks(db *sql.DB) http.HandlerFunc {
	var tasks []models.Task
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, sequence, priority, user_id FROM tasks")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for rows.Next() {
			var task models.Task
			if err := rows.Scan(&task.Id, &task.Sequence, &task.UserId, &task.Priority); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			tasks = append(tasks, task)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	}
}

func CreateTask(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var task models.Task
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		err = db.QueryRow("INSERT INTO tasks(sequence, user_id, priority) VALUES($1, $2, $3) RETURNING id", task.Sequence, task.UserId, task.Priority).Scan(&task.Id)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	}
}

func UpdateTask(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func DeleteTask(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
