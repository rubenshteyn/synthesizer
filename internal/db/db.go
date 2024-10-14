package db

import (
	"database/sql"
	"log"
)

func Connect() *sql.DB {
	var err error
	connStr := "user=postgres dbname=synthesizer sslmode=disable password=1234"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
