package model

import (
	"database/sql"
	"log"
)

func Connect() *sql.DB {
	db, err := sql.Open("postgres", "awtpbzyctlpydm:a7bf40c39496f73a03e7412befbc787d29138445d7fce2a34bf31df40cf07d96@(tcp:localhost:5432)")
	if err != nil {
		log.Fatal(err)
	}
	return db
}