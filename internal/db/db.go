package db

import (
	"database/sql"
	"log"
)

func Connect() *sql.DB {
	var err error
	connStr := "user=synthesizer_admin dbname=synthesizer sslmode=disable password=synthesizer_admin_password"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
