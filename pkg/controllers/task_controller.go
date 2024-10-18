package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"synthesizer/pkg/models"

	"github.com/gorilla/mux"
)

func GetTasks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tx, err := db.Begin()
		if err != nil {
			http.Error(w, "bad_request", http.StatusBadRequest)
			return
		}

		var tasks []models.Task
		vars := mux.Vars(r)
		id := vars["id"]

		rows, err := tx.Query("SELECT id, sequence, priority, user_id FROM tasks WHERE user_id = $1", id)
		if err != nil {
			http.Error(w, "update", http.StatusBadRequest)
			return
		}

		for rows.Next() {
			var task models.Task
			if err := rows.Scan(&task.Id, &task.Sequence, &task.Priority, &task.UserId); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			tasks = append(tasks, task)
		}
		rows.Close()

		defer func() {
			if err != nil {
				tx.Rollback()
				return
			}
			err = tx.Commit()
		}()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	}
}

func CreateTask(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tx, err := db.Begin()
		if err != nil {
			http.Error(w, "bad_request", http.StatusBadRequest)
			return
		}

		var task models.Task
		err = json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, "bad_request", http.StatusBadRequest)
			return
		}

		_, err = tx.Exec("INSERT INTO tasks(sequence, user_id, priority, status) VALUES($1, $2, $3, $4) RETURNING id", task.Sequence, task.UserId, task.Priority, task.Status)
		if err != nil {
			http.Error(w, "update", http.StatusBadRequest)
			return
		}

		defer func() {
			if err != nil {
				tx.Rollback()
				return
			}
			err = tx.Commit()
		}()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	}
}

func UpdateTaskStatus(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tx, err := db.Begin()
		if err != nil {
			http.Error(w, "bad_request", http.StatusBadRequest)
			return
		}

		var task models.Task
		err = json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, "bad_request", http.StatusBadRequest)
			return
		}

		_, err = tx.Exec("UPDATE tasks SET status = $1 WHERE id = $2", task.Status, task.Id)
		if err != nil {
			http.Error(w, "update", http.StatusBadRequest)
			return
		}

		defer func() {
			if err != nil {
				tx.Rollback()
				return
			}
			err = tx.Commit()
		}()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	}
}

func UpdateTaskSequence(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tx, err := db.Begin()
		if err != nil {
			http.Error(w, "bad_request", http.StatusBadRequest)
			return
		}

		var task models.Task
		err = json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, "bad_request", http.StatusBadRequest)
			return
		}

		_, err = tx.Exec("UPDATE tasks SET sequence = $1 WHERE id = $2", task.Sequence, task.Id)
		if err != nil {
			http.Error(w, "update", http.StatusBadRequest)
			return
		}

		defer func() {
			if err != nil {
				tx.Rollback()
				return
			}
			err = tx.Commit()
		}()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	}
}

func DeleteTask(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tx, err := db.Begin()
		if err != nil {
			http.Error(w, "bad_request", http.StatusBadRequest)
			return
		}

		vars := mux.Vars(r)
		id := vars["id"]

		_, err = tx.Exec("DELETE FROM tasks WHERE = $1", id)
		if err != nil {
			http.Error(w, "internal", http.StatusInternalServerError)
			return
		}

		defer func() {
			if err != nil {
				tx.Rollback()
				return
			}
			err = tx.Commit()
		}()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(id)
	}
}
