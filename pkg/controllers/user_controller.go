package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"synthesizer/pkg/models"
)

func GetUsers(db *sql.DB) http.HandlerFunc {
	var users []models.User
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name, password, task_id FROM users")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for rows.Next() {
			var user models.User
			if err := rows.Scan(&user.Id, &user.Name, &user.Password, &user.TaskId); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			users = append(users, user)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		var name int
		getUserQuery := "SELECT user FROM users WHERE name = $1"
		userIsExist := db.QueryRow(getUserQuery, user.Name).Scan(&name)
		if userIsExist != nil {
			http.Error(w, "User is exist", http.StatusNotFound)
			return
		}

		err = db.QueryRow("INSERT INTO users(name, task_id, password) VALUES($1, $2, $3) RETURNING id", user.Name, user.TaskId, user.Password).Scan(&user.Id)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

func UpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
