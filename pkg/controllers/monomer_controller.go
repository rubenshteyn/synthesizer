package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"synthesizer/pkg/models"

	"github.com/gorilla/mux"
)

func GetMonomers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tx, err := db.Begin()
		if err != nil {
			http.Error(w, "bad_request", http.StatusBadRequest)
			return
		}

		var monomers []models.Monomer
		rows, err := tx.Query("SELECT name, count FROM monomers")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for rows.Next() {
			var monomer models.Monomer
			if err := rows.Scan(&monomer.Name, &monomer.Count); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			monomers = append(monomers, monomer)
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
		json.NewEncoder(w).Encode(monomers)
	}
}

func GetMonomerCount(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tx, err := db.Begin()
		if err != nil {
			http.Error(w, "bad_request", http.StatusBadRequest)
			return
		}

		vars := mux.Vars(r)
		name := vars["name"]
		count, err := tx.Query("SELECT count FROM monomers WHERE name = $1", name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
		json.NewEncoder(w).Encode(count)
	}
}

func CreateMonomer(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tx, err := db.Begin()
		if err != nil {
			http.Error(w, "bad_request", http.StatusBadRequest)
			return
		}

		var monomer models.Monomer
		err = json.NewDecoder(r.Body).Decode(&monomer)
		if err != nil {
			http.Error(w, "bad_request", http.StatusBadRequest)
			return
		}

		var name string
		exist := tx.QueryRow("SELECT name FROM monomers WHERE name = $1", monomer.Name).Scan(&name)
		if exist == nil {
			http.Error(w, "is_exist", http.StatusBadRequest)
			return
		}

		err = tx.QueryRow("INSERT INTO monomers(name, count) VALUES($1, $2) RETURNING name", monomer.Name, monomer.Count).Scan(&name)
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
		json.NewEncoder(w).Encode(monomer)
	}
}

func UpdateMonomerCount(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tx, err := db.Begin()
		if err != nil {
			http.Error(w, "bad_request", http.StatusBadRequest)
			return
		}

		var monomer models.Monomer
		err = json.NewDecoder(r.Body).Decode(&monomer)
		if err != nil {
			http.Error(w, "bad_request", http.StatusBadRequest)
			return
		}

		_, err = tx.Exec("UPDATE monomers SET count = $1 WHERE name = $2", monomer.Count, monomer.Name)
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
		json.NewEncoder(w).Encode(monomer)
	}
}
