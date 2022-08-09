package model

import (
	"database/sql"
	"log"
	"fmt"
)

const (
	host = "localhost"
	port = 5432
	user = "awtpbzyctlpydm"
	pass = "a7bf40c39496f73a03e7412befbc787d29138445d7fce2a34bf31df40cf07d96"
	dbname = "d4ehughfapgq0k"

)

func Connect() *sql.DB {
	db, err := sql.Open("postgres", "awtpbzyctlpydm:a7bf40c39496f73a03e7412befbc787d29138445d7fce2a34bf31df40cf07d96@(tcp:localhost:5432)")
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Connected to postgres")
	return db

	
}