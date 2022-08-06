package model

import (
	"database/sql"
	"log"
)

func Connect() *sql.DB {
	db, err := sql.Open("postgres", "root:1234@(tcp:localhost:6379)")
	if err != nil {
		log.Fatal(err)
	}
	return db
}