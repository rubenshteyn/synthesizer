package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"synthesizer/pkg/models"

	"golang.org/x/crypto/bcrypt"
)

// Хеширование пароля
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Проверка пароля
func checkPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GetUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var users []models.User
		rows, err := db.Query("SELECT id, name, password FROM users")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for rows.Next() {
			var user models.User
			if err := rows.Scan(&user.Id, &user.Name, &user.Password); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			users = append(users, user)
		}
		rows.Close()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

func Authentication(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "bad_request", http.StatusBadRequest)
			return
		}

		exist := db.QueryRow("SELECT name FROM users WHERE name = $1", user.Name).Scan(&user.Name)
		if exist == nil {
			http.Error(w, "is_exist", http.StatusBadRequest)
			return
		}

		hashedPassword, err := hashPassword(user.Password)
		if err != nil {
			log.Fatalf("Hashed password error %v", err)
		}

		err = db.QueryRow("INSERT INTO users(name, password) VALUES($1, $2) RETURNING id", user.Name, hashedPassword).Scan(&user.Id)
		if err != nil {
			http.Error(w, "internal", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "bad_request", http.StatusBadRequest)
			return
		}

		var hashedPassword string
		err = db.QueryRow("SELECT password FROM users WHERE name=$1", user.Name).Scan(&hashedPassword)
		if err != nil {
			http.Error(w, "not_found", http.StatusNotFound)
			return
		}

		err = checkPassword(hashedPassword, user.Password)
		if err != nil {
			http.Error(w, "incorrect_password", http.StatusBadRequest)
			return
		}

		isAlreadyAuth := IsAuthenticated(r)
		fmt.Println(isAlreadyAuth)
		if isAlreadyAuth != true {
			err = CreateSession(w, r, user.Name)
			if err != nil {
				http.Error(w, "bad_request", http.StatusBadRequest)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(true)
	}
}

func UpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
