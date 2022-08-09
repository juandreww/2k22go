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
	password = "a7bf40c39496f73a03e7412befbc787d29138445d7fce2a34bf31df40cf07d96"
	dbname = "d4ehughfapgq0k"

)

func Connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Connected to postgres")
	return db

	
}